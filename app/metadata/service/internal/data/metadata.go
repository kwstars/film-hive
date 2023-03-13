package data

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/cache/v9"
	v1 "github.com/kwstars/film-hive/api/metadata/service/v1"
	"github.com/kwstars/film-hive/app/metadata/service/internal/biz"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type metadataRepo struct {
	data *Data
	log  *log.Helper
}

func (m *metadataRepo) GetMetadata(ctx context.Context, id uint64) (data *biz.Metadata, err error) {
	metadata := &Metadata{}
	data = &biz.Metadata{}
	key := MetadataStringKey + ":" + strconv.FormatUint(id, 10)
	if err = m.data.cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: data,
		TTL:   RedisDefaultExpire,
		SetNX: true,
		Do: func(*cache.Item) (interface{}, error) {
			err = m.data.db.WithContext(ctx).First(metadata, id).Error
			if err != nil {
				switch {
				case errors.Is(err, gorm.ErrRecordNotFound):
					return nil, v1.ErrorMetadataNotFound("metadata was not found, id: %d", id)
				default:
					return nil, errors.Wrapf(err, "GetMetadata failed, id: %d", id)
				}
			}
			data = &biz.Metadata{
				ID:          metadata.ID,
				Title:       metadata.Title,
				Description: metadata.Description,
				Director:    metadata.Director,
			}
			return data, nil
		},
	}); err != nil {
		return nil, errors.Wrapf(err, "GetMetadata, id: %d", id)
	}
	return
}

// NewMetadataRepo .
func NewMetadataRepo(data *Data, logger log.Logger) biz.MetadataRepo {
	return &metadataRepo{
		data: data,
		log:  log.NewHelper(logger, log.WithMessageKey("metadata repo")),
	}
}
