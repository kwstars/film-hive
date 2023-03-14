package data

import "time"

var Tables = []any{
	new(Rating),
}

const (
	RatingStringKey = "rat"
)

const (
	RedisDefaultExpire = 8 * time.Hour
)

// Rating defines the movie rating.
type Rating struct {
	RecordID   uint64 `gorm:"primaryKey"`
	RecordType uint64 `gorm:"primaryKey;comment:'评论类型1电影2电视剧'"`
	UserID     uint64 `gorm:"primaryKey;comment:'用户ID'"`
	Value      uint32 `gorm:"not null;comment:'评分'"`
}
