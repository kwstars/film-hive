package data

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/cache/v9"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/google/wire"
	"github.com/kwstars/film-hive/app/metadata/service/internal/conf"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMetadataRepo)

// Data .
type Data struct {
	db       *gorm.DB
	cache    *cache.Cache
	dlock    *redsync.Redsync
	dataConf *conf.Data
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger) (data *Data, closer func(), err error) {
	var (
		helper = log.NewHelper(logger, log.WithMessageKey("metadata data"))
		db     *gorm.DB
		rdb    *redis.Client
	)

	if rdb, err = newRedis(c.Data); err != nil {
		return
	}
	pool := goredis.NewPool(rdb)
	mycache := cache.New(&cache.Options{
		Redis: rdb,
		// LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	if db, err = newDB(c.Data, helper); err != nil {
		return
	}

	closer = func() {
		helper.Info("closing the data resources")
		if err = rdb.Close(); err != nil {
			helper.Infof("Redis close: %v", err)
		}
		var sqlDB *sql.DB
		if sqlDB, err = db.DB(); err != nil {
			err = sqlDB.Close()
			helper.Infof("MySQL close: %v", err)
		}
	}
	data = &Data{
		db:       db,
		cache:    mycache,
		dlock:    redsync.New(pool),
		dataConf: c.Data,
	}
	return
}

func newRedis(c *conf.Data) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	// Enable tracing instrumentation.
	if err = redisotel.InstrumentTracing(rdb); err != nil {
		return nil, err
	}

	return
}

func newDB(c *conf.Data, helper *log.Helper) (db *gorm.DB, err error) {
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Database.GetSource(), // data source name
		DefaultStringSize:         256,                    // default size for string fields
		DisableDatetimePrecision:  true,                   // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                   // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                   // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                  // autoconfigure based on currently MySQL version
	}), &gorm.Config{}); err != nil {
		return
	}
	if err = db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		return
	}
	// NOTE: GORM has problem detecting existing columns, see
	// https://github.com/gogs/gogs/issues/6091. Therefore, only use it to create new
	// tables, and do customize migration with future changes.
	for _, table := range Tables {
		if db.Migrator().HasTable(table) {
			continue
		}
		name := strings.TrimPrefix(fmt.Sprintf("%T", table), "*db.")
		if err = db.Migrator().AutoMigrate(table); err != nil {
			return nil, errors.Wrapf(err, "auto migrate %q", name)
		}
		helper.Infof("Auto migrated %q", name)
	}
	return
}
