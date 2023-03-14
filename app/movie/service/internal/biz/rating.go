package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	metadata "github.com/kwstars/film-hive/api/metadata/service/v1"
	rating "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/kwstars/film-hive/pkg/sync/errgroup"
)

type MovieRepo interface {
	ListMovies(ctx context.Context, recordType, recordID string) (movies []Movie, err error)
	CreateMovie(ctx context.Context, recordType, recordID string, movie *Movie) (err error)
}

type MovieUsecase struct {
	repo MovieRepo
	log  *log.Helper
	mc   metadata.MetadataServiceClient
	rc   rating.RatingServiceClient
}

func NewMovieUsecase(mc metadata.MetadataServiceClient, rc rating.RatingServiceClient, repo MovieRepo, logger log.Logger) *MovieUsecase {
	return &MovieUsecase{
		repo: repo,
		log:  log.NewHelper(logger, log.WithMessageKey("movie biz")),
		mc:   mc,
		rc:   rc,
	}
}

func (m *MovieUsecase) GetMovieDetail(ctx context.Context, ID uint64) (resp *MovieDetail, err error) {
	var (
		ratingResp   *rating.GetAggregatedRatingResponse
		metadataResp *metadata.GetMetadataResponse
	)

	eg := errgroup.WithContext(ctx)
	eg.Go(func(ctx context.Context) error {
		if ratingResp, err = m.rc.GetAggregatedRating(ctx, &rating.GetAggregatedRatingRequest{RecordType: 1, RecordId: ID}); err != nil {
			log.Errorf("rating service failed, id: %d, %v", ID, err)
			return err
		}
		return nil
	})
	eg.Go(func(ctx context.Context) error {
		if metadataResp, err = m.mc.GetMetadata(ctx, &metadata.GetMetadataRequest{Id: ID}); err != nil {
			log.Errorf("metadata service failed, id: %d, %v", ID, err)
			return err
		}
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	return &MovieDetail{
		Rating: ratingResp.GetAvgRating(),
		Metadata: &Metadata{
			Id:          metadataResp.GetId(),
			Title:       metadataResp.GetTitle(),
			Description: metadataResp.GetDescription(),
			Director:    metadataResp.GetDirector(),
		},
	}, nil
}
