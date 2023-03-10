package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/pkg/errors"
)

type RatingRepo interface {
	ListRatings(ctx context.Context, recordType, recordID string) (ratings []Rating, err error)
	CreateRating(ctx context.Context, recordType, recordID string, rating *Rating) (err error)
}

type Rating struct {
	RecordID   string `json:"recordId"`
	RecordType string `json:"recordType"`
	UserID     string `json:"userId"`
	Value      uint32 `json:"value"`
}

type RatingUsecase struct {
	repo RatingRepo
	log  *log.Helper
}

func NewRatingUsecase(repo RatingRepo, logger log.Logger) *RatingUsecase {
	return &RatingUsecase{
		repo: repo,
		log:  log.NewHelper(logger, log.WithMessageKey("rating")),
	}
}

// GetAggregatedRating 获取平局的评分
func (c *RatingUsecase) GetAggregatedRating(ctx context.Context, recordType, recordID string) (ar float64, err error) {
	ratings, err := c.repo.ListRatings(ctx, recordType, recordID)
	if err != nil {
		switch {
		case v1.IsRatingNotFound(err):
			c.log.Errorf("rating not found %v", err)
			return
		default:
			c.log.Errorf("ListRating err: %v", err)
			return 0, errors.New("unknown error")
		}
	}
	sum := float64(0)
	for _, r := range ratings {
		sum += float64(r.Value)
	}
	return sum / float64(len(ratings)), nil
}

// CreateRating 添加一个评分
func (c *RatingUsecase) CreateRating(ctx context.Context, recordType, recordID string, rating *Rating) error {
	return c.repo.CreateRating(ctx, recordType, recordID, rating)
}
