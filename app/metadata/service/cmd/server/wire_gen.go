// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/kwstars/film-hive/app/metadata/service/internal/biz"
	"github.com/kwstars/film-hive/app/metadata/service/internal/conf"
	"github.com/kwstars/film-hive/app/metadata/service/internal/data"
	"github.com/kwstars/film-hive/app/metadata/service/internal/server"
	"github.com/kwstars/film-hive/app/metadata/service/internal/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(bootstrap *conf.Bootstrap, tp *trace.TracerProvider) (*kratos.App, func(), error) {
	logger, cleanup := server.NewLogger()
	registrar, err := server.NewNamingClient(bootstrap)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dataData, cleanup2, err := data.NewData(bootstrap, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	metadataRepo := data.NewMetadataRepo(dataData, logger)
	metadataUsecase := biz.NewMetadataUsecase(metadataRepo, logger)
	metadataService := service.NewMetadataService(metadataUsecase)
	grpcServer := server.NewGRPCServer(bootstrap, metadataService, logger)
	httpServer := server.NewHTTPServer(bootstrap, metadataService, logger)
	app := newApp(logger, registrar, grpcServer, httpServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
