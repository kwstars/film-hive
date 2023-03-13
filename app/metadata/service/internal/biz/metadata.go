package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kwstars/film-hive/api/metadata/service/v1"
	"github.com/pkg/errors"
)

//go:generate mockgen -package biz -destination metadata_mock.go -source metadata.go MetadataRepo

type MetadataRepo interface {
	GetMetadata(ctx context.Context, id uint64) (data *Metadata, err error)
}

type Metadata struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
}

type MetadataUsecase struct {
	repo MetadataRepo
	log  *log.Helper
}

func NewMetadataUsecase(repo MetadataRepo, logger log.Logger) *MetadataUsecase {
	return &MetadataUsecase{
		repo: repo,
		log:  log.NewHelper(logger, log.WithMessageKey("metadata data")),
	}
}

func (mu *MetadataUsecase) GetMetadata(ctx context.Context, id uint64) (metadata *Metadata, err error) {
	data, err := mu.repo.GetMetadata(ctx, id)
	if err != nil {
		switch {
		case v1.IsMetadataNotFound(err):
			mu.log.Errorf("metadata not found %v", err)
			return
		default:
			mu.log.Errorf("GetMetadata err: %v", err)
			return nil, errors.New("unknown error")
		}
	}
	return data, nil
}
