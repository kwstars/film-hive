package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kwstars/film-hive/api/rating/service/v1"
)

type RatingRepo interface {
	ListRatings(ctx context.Context, recordType, recordID uint64) (ratings []uint32, err error)
	CreateRating(ctx context.Context, uid uint64, rating *Rating) (err error)
}

type Rating struct {
	RecordID   uint64 `json:"recordId"`
	RecordType uint64 `json:"recordType"`
	UserID     uint64 `json:"userId"`
	Value      uint32 `json:"value"`
}

type RatingUsecase struct {
	repo RatingRepo
	log  *log.Helper
}

func NewRatingUsecase(repo RatingRepo, logger log.Logger) *RatingUsecase {
	return &RatingUsecase{
		repo: repo,
		log:  log.NewHelper(logger, log.WithMessageKey("rating biz")),
	}
}

// GetAggregatedRating 获取平局的评分
func (c *RatingUsecase) GetAggregatedRating(ctx context.Context, recordType, recordID uint64) (ar float64, err error) {
	ratingValues, err := c.repo.ListRatings(ctx, recordType, recordID)
	if err != nil {
		switch {
		case v1.IsRatingNotFound(err):
			c.log.Errorf("rating not found %v", err)
			return
		default:
			c.log.Errorf("ListRating err: %v", err)
			return 0, errors.New(errors.UnknownCode, errors.UnknownReason, "服务端错误")
		}
	}
	sum := float64(0)
	for _, value := range ratingValues {
		sum += float64(value)
	}
	return sum / float64(len(ratingValues)), nil
}

// CreateRating 添加一个评分
func (c *RatingUsecase) CreateRating(ctx context.Context, uid uint64, rating *Rating) error {
	if err := c.repo.CreateRating(ctx, uid, rating); err != nil {
		c.log.Errorf("CreateRating failed, uid: %d, rating: %+v, err: %v", uid, rating, err)
		return errors.New(errors.UnknownCode, errors.UnknownReason, "服务端错误")
	}
	return nil
}
