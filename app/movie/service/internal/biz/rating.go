package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kwstars/film-hive/api/movie/service/v1"
	"github.com/pkg/errors"
)

type MovieRepo interface {
	ListMovies(ctx context.Context, recordType, recordID string) (movies []Movie, err error)
	CreateMovie(ctx context.Context, recordType, recordID string, movie *Movie) (err error)
}

type Movie struct {
	RecordID   string `json:"recordId"`
	RecordType string `json:"recordType"`
	UserID     string `json:"userId"`
	Value      uint32 `json:"value"`
}

type MovieUsecase struct {
	repo MovieRepo
	log  *log.Helper
}

func NewMovieUsecase(repo MovieRepo, logger log.Logger) *MovieUsecase {
	return &MovieUsecase{
		repo: repo,
		log:  log.NewHelper(logger, log.WithMessageKey("movie biz")),
	}
}

// GetAggregatedMovie 获取平局的评分
func (c *MovieUsecase) GetAggregatedMovie(ctx context.Context, recordType, recordID string) (ar float64, err error) {
	movies, err := c.repo.ListMovies(ctx, recordType, recordID)
	if err != nil {
		switch {
		case v1.IsMovieNotFound(err):
			c.log.Errorf("GetAggregatedMovie movie not found: %v", err)
			return
		default:
			c.log.Errorf("ListMovie err: %v", err)
			return 0, errors.New("unknown error")
		}
	}
	sum := float64(0)
	for _, r := range movies {
		sum += float64(r.Value)
	}
	return sum / float64(len(movies)), nil
}

// CreateMovie 添加一个评分
func (c *MovieUsecase) CreateMovie(ctx context.Context, recordType, recordID string, movie *Movie) error {
	return c.repo.CreateMovie(ctx, recordType, recordID, movie)
}
