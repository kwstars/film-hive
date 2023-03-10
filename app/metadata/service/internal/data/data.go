package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/kwstars/film-hive/app/metadata/service/internal/conf"
	"sync"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMetadataRepo)

// Data .
type Data struct {
	mutex    sync.RWMutex
	metaData map[string]*Metadata
}

// NewData .
func NewData(c *conf.Bootstrap, _ log.Logger) (data *Data, closer func(), err error) {

	d := map[string]*Metadata{
		"111111": {"11111", "11111", "11111", "11111"},
		"222222": {"222222", "222222", "222222", "222222"},
	}

	return &Data{
		metaData: d,
	}, func() {}, nil
}
