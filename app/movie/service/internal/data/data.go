package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/kwstars/film-hive/app/movie/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMovieRepo)

// Data .
type Data struct {
	movie map[string]map[string][]Movie // map[recordType]map[recordID][]Movie
}

// NewData .
func NewData(_ *conf.Bootstrap, _ log.Logger) (data *Data, closer func(), err error) {
	d := map[string]map[string][]Movie{
		"1": {"1": {{RecordID: "1", RecordType: "1", UserID: "1", Value: 1}, {RecordID: "3", RecordType: "3", UserID: "3", Value: 3}}},
		"2": {"2": {{RecordID: "2", RecordType: "2", UserID: "2", Value: 2}}},
		"3": {"3": {{RecordID: "3", RecordType: "3", UserID: "3", Value: 3}}},
	}
	return &Data{
		movie: d,
	}, func() {}, nil
}
