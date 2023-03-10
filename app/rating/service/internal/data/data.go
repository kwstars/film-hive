package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/kwstars/film-hive/app/rating/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRatingRepo)

// Data .
type Data struct {
	rating map[string]map[string][]Rating // map[recordType]map[recordID][]Rating
}

// NewData .
func NewData(c *conf.Bootstrap, _ log.Logger) (data *Data, closer func(), err error) {
	d := map[string]map[string][]Rating{
		"1": {"1": {{RecordID: "1", RecordType: "1", UserID: "1", Value: 1}, {RecordID: "3", RecordType: "3", UserID: "3", Value: 3}}},
		"2": {"2": {{RecordID: "2", RecordType: "2", UserID: "2", Value: 2}}},
		"3": {"3": {{RecordID: "3", RecordType: "3", UserID: "3", Value: 3}}},
	}
	return &Data{
		rating: d,
	}, func() {}, nil
}
