//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/kwstars/film-hive/app/rating/service/internal/biz"
	"github.com/kwstars/film-hive/app/rating/service/internal/conf"
	"github.com/kwstars/film-hive/app/rating/service/internal/data"
	"github.com/kwstars/film-hive/app/rating/service/internal/server"
	"github.com/kwstars/film-hive/app/rating/service/internal/service"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2"
)

// initApp init kratos application.
func initApp(bootstrap *conf.Bootstrap, tp *trace.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
