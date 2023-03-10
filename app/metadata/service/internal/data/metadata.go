package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kwstars/film-hive/api/metadata/service/v1"
	"github.com/kwstars/film-hive/app/metadata/service/internal/biz"
)


type metadataRepo struct {
	data *Data
	log  *log.Helper
}


func (a *metadataRepo) GetMetadata(ctx context.Context, id string) (data *biz.Metadata, err error) {
	a.data.mutex.RLock()
	defer a.data.mutex.RUnlock()
	if metadata, ok := a.data.metaData[id]; ok {
		d := &biz.Metadata{
			ID:          metadata.ID,
			Title:       metadata.Title,
			Description: metadata.Description,
			Director:    metadata.Director,
		}
		return d, nil
	}
	return nil, v1.ErrorMetadataNotFound("metadata was not found, id: %s", id)
}

// NewMetadataRepo .
func NewMetadataRepo(data *Data, logger log.Logger) biz.MetadataRepo {
	return &metadataRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
