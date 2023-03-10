package service

import (
	"context"
	metadata "github.com/kwstars/film-hive/api/metadata/service/v1"
	v1 "github.com/kwstars/film-hive/api/movie/service/v1"
	rating "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/kwstars/film-hive/app/movie/service/internal/biz"
	"github.com/kwstars/film-hive/app/movie/service/internal/conf"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ v1.MovieServiceServer = (*MovieService)(nil)

type MovieService struct {
	v1.UnimplementedMovieServiceServer
	uc       *biz.MovieUsecase
	rating   rating.RatingServiceClient
	metadata metadata.MetadataServiceClient
}

func NewMovieService(c *conf.Bootstrap, rc naming_client.INamingClient, uc *biz.MovieUsecase) *MovieService {
	return &MovieService{
		uc:       uc,
		rating:   newRatingClient(c, rc),
		metadata: newMetadataClient(c, rc),
	}
}

func (m *MovieService) GetMovieDetail(ctx context.Context, req *v1.GetMovieDetailRequest) (resp *v1.GetMovieDetailResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieDetail not implemented")
}

func (m *MovieService) GetAggregatedRating(ctx context.Context, in *rating.GetAggregatedRatingRequest) (*rating.GetAggregatedRatingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieService) CreateRating(ctx context.Context, in *rating.CreateRatingRequest) (*rating.CreateRatingResponse, error) {
	//TODO implement me
	panic("implement me")
}
