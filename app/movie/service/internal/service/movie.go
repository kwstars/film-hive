package service

import (
	"context"

	metadata "github.com/kwstars/film-hive/api/metadata/service/v1"
	v1 "github.com/kwstars/film-hive/api/movie/service/v1"
	rating "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/kwstars/film-hive/app/movie/service/internal/biz"
	"github.com/kwstars/film-hive/pkg/sync/errgroup"
)

type MovieService struct {
	v1.UnimplementedMovieServiceServer
	uc       *biz.MovieUsecase
	rating   rating.RatingServiceClient
	metadata metadata.MetadataServiceClient
}

func NewMovieService(rc rating.RatingServiceClient, mc metadata.MetadataServiceClient, uc *biz.MovieUsecase) *MovieService {
	return &MovieService{
		uc:       uc,
		rating:   rc,
		metadata: mc,
	}
}

func (m *MovieService) GetMovieDetail(ctx context.Context, req *v1.GetMovieDetailRequest) (resp *v1.GetMovieDetailResponse, err error) {
	var (
		ratingResp   *rating.GetAggregatedRatingResponse
		meatdataResp *metadata.GetMetadataResponse
	)

	eg := errgroup.WithContext(ctx)
	eg.Go(func(ctx context.Context) error {
		if ratingResp, err = m.rating.GetAggregatedRating(ctx, &rating.GetAggregatedRatingRequest{RecordType: 1, RecordId: req.GetId()}); err != nil {
			return err
		}
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		if meatdataResp, err = m.metadata.GetMetadata(ctx, &metadata.GetMetadataRequest{Id: req.GetId()}); err != nil {
			return err
		}
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &v1.GetMovieDetailResponse{
		Rating: ratingResp.GetAvgRating(),
		Metadata: &v1.GetMovieDetailResponse_Metadata{
			Id:          meatdataResp.GetId(),
			Title:       meatdataResp.GetTitle(),
			Description: meatdataResp.GetDescription(),
			Director:    meatdataResp.GetDirector(),
		},
	}
	return
}
