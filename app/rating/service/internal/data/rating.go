package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/kwstars/film-hive/app/rating/service/internal/biz"
)

type ratingRepo struct {
	data *Data
	log  *log.Helper
}

func (r *ratingRepo) ListRatings(ctx context.Context, recordType, recordID string) (ratings []biz.Rating, err error) {
	var (
		ok bool
		rs []Rating
	)
	if _, ok = r.data.rating[recordType]; !ok {
		return nil, v1.ErrorRatingNotFound("no record of record_type(%s)", recordType)
	}
	if rs, ok = r.data.rating[recordType][recordID]; !ok || len(rs) == 0 {
		return nil, v1.ErrorRatingNotFound("no record of record_type(%s) and record_id(%s)", recordType, recordID)
	}

	ratings = make([]biz.Rating, 0, len(rs))
	for _, rating := range rs {
		t := biz.Rating{
			RecordID:   rating.RecordID,
			RecordType: rating.RecordType,
			UserID:     rating.UserID,
			Value:      rating.Value,
		}
		ratings = append(ratings, t)
	}
	return
}

func (r *ratingRepo) CreateRating(ctx context.Context, recordType, recordID string, rating *biz.Rating) (err error) {
	t := Rating{
		RecordID:   rating.RecordID,
		RecordType: rating.RecordType,
		UserID:     rating.UserID,
		Value:      rating.Value,
	}
	if _, ok := r.data.rating[recordType]; !ok {
		r.data.rating[recordType] = map[string][]Rating{}
	}
	r.data.rating[recordType][recordID] = append(r.data.rating[recordType][recordID], t)
	return
}

// NewRatingRepo .
func NewRatingRepo(data *Data, logger log.Logger) biz.RatingRepo {
	return &ratingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
