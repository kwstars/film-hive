package data

import "time"

var Tables = []any{
	new(Metadata),
}

const (
	RedisDefaultExpire = 8 * time.Hour
)

const (
	MetadataStringKey string = "md"
)

// Metadata defines the movie metadata.
type Metadata struct {
	ID          uint64 `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(30);not null;comment:'名字'"`
	Description string `gorm:"type:varchar(255);not null;comment:'描述'"`
	Director    string `gorm:"type:varchar(50);not null;comment:'导演'"`
}
