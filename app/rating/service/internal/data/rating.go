package data

import (
	"context"
	"github.com/go-redis/cache/v9"
	"github.com/golang/protobuf/proto"
	v1 "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/kwstars/film-hive/app/rating/service/internal/biz"
)

type ratingRepo struct {
	data *Data
	log  *log.Helper
}

// ListRatings TODO
func (r *ratingRepo) ListRatings(ctx context.Context, recordType, recordID uint64) (rs []uint32, err error) {
	key := RatingStringKey + ":" + strconv.FormatUint(recordType, 10) + ":" + strconv.FormatUint(recordID, 10)
	var ratings = []Rating{}
	rs = []uint32{}
	if err = r.data.cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: &rs,
		TTL:   RedisDefaultExpire,
		Do: func(item *cache.Item) (interface{}, error) {
			err = r.data.db.Select("value").Find(&ratings, "record_id = ? AND record_type = ?", recordID, recordType).Error
			if err != nil {
				switch {
				case errors.Is(err, gorm.ErrRecordNotFound):
					return nil, v1.ErrorRatingNotFound("rating was not found, type: %d, id: %d", recordType, recordID)
				default:
					return nil, errors.Wrapf(err, "ListRatings failed")
				}
			}
			rs = make([]uint32, 0, len(ratings))
			for _, v := range ratings {
				rs = append(rs, v.Value)
			}
			return rs, nil
		},
		SetNX: true,
	}); err != nil {
		return nil, errors.Wrap(err, "ListRatings failed")
	}
	return
}

// CreateRating 创建评分.
func (r *ratingRepo) CreateRating(ctx context.Context, uid uint64, rating *biz.Rating) (err error) {
	key := strconv.FormatUint(uid, 10)
	tmp := &v1.Rating{
		RecordType:  v1.RECORDTYPE(rating.RecordType),
		RecordId:    rating.RecordID,
		UserId:      rating.UserID,
		RatingValue: rating.Value,
	}
	msg, err := proto.Marshal(tmp)
	if err != nil {
		return errors.Wrap(err, "CreateRating")
	}
	return r.data.PushMsg(key, msg)
}

// NewRatingRepo .
func NewRatingRepo(data *Data, logger log.Logger) biz.RatingRepo {
	return &ratingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
