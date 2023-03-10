package service

import (
	"context"

	v1 "github.com/kwstars/film-hive/api/metadata/service/v1"
	"github.com/kwstars/film-hive/app/metadata/service/internal/biz"
)

var _ v1.MetadataServiceServer = (*MetadataService)(nil)

type MetadataService struct {
	v1.UnimplementedMetadataServiceServer
	uc *biz.MetadataUsecase
}

func NewMetadataService(uc *biz.MetadataUsecase) *MetadataService {
	return &MetadataService{uc: uc}
}

func (m *MetadataService) GetMetadata(ctx context.Context, req *v1.GetMetadataRequest) (resp *v1.GetMetadataResponse, err error) {
	metadata, err := m.uc.GetMetadata(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	resp = &v1.GetMetadataResponse{
		Id:          metadata.ID,
		Title:       metadata.Title,
		Description: metadata.Description,
		Director:    metadata.Director,
	}

	return
}
