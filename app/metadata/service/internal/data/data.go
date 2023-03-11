package data

import (
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/kwstars/film-hive/app/metadata/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMetadataRepo)

// Data .
type Data struct {
	mutex    sync.RWMutex
	metaData map[string]*Metadata
}

// NewData .
func NewData(c *conf.Bootstrap, _ log.Logger) (data *Data, closer func(), err error) {
	d := map[string]*Metadata{
		"1": {"1", "1", "1", "1"},
		"2": {"2", "2", "2", "2"},
	}

	return &Data{
		metaData: d,
	}, func() {}, nil
}
