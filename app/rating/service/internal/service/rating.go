package service

import (
	"context"
	v1 "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/kwstars/film-hive/app/rating/service/internal/biz"
)

var _ v1.RatingServiceServer = (*RatingService)(nil)

type RatingService struct {
	v1.UnimplementedRatingServiceServer
	uc *biz.RatingUsecase
}

func NewRatingService(uc *biz.RatingUsecase) *RatingService {
	return &RatingService{uc: uc}
}

func (r *RatingService) GetAggregatedRating(ctx context.Context, req *v1.GetAggregatedRatingRequest) (resp *v1.GetAggregatedRatingResponse, err error) {
	ar, err := r.uc.GetAggregatedRating(ctx, req.GetRecordType(), req.GetRecordId())
	if err != nil {
		return
	}
	resp = &v1.GetAggregatedRatingResponse{
		AvgRating: ar,
	}
	return
}

func (r *RatingService) CreateRating(ctx context.Context, req *v1.CreateRatingRequest) (resp *v1.CreateRatingResponse, err error) {
	t := &biz.Rating{
		RecordID:   req.GetRecordId(),
		RecordType: req.GetRecordType(),
		UserID:     req.GetUserId(),
		Value:      req.GetRatingValue(),
	}
	return &v1.CreateRatingResponse{}, r.uc.CreateRating(ctx, req.GetRecordType(), req.GetUserId(), t)
}
