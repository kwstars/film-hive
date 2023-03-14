//go:build wireinject
// +build wireinject

package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/kwstars/film-hive/app/rating/service/internal/biz"
	"github.com/kwstars/film-hive/app/rating/service/internal/conf"
	"github.com/kwstars/film-hive/app/rating/service/internal/data"
)

func createRatingService(bootstrap *conf.Bootstrap, log log.Logger) (*RatingService, func(), error) {
	panic(wire.Build(ProviderSet, data.ProviderSet, biz.ProviderSet))
}
