package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	metadata "github.com/kwstars/film-hive/api/metadata/service/v1"
	"github.com/kwstars/film-hive/app/movie/service/internal/conf"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

func NewMetadataGRPCClient(c *conf.Bootstrap, rc naming_client.INamingClient) (metadata.MetadataServiceClient, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("discovery:///metadata.grpc"),
		grpc.WithDiscovery(nacos.New(rc)),
		grpc.WithMiddleware(
			tracing.Client(),
		),
	)
	if err != nil {
		return nil, func() {}, err
	}
	return metadata.NewMetadataServiceClient(conn), func() {
		if err = conn.Close(); err != nil {
			log.Errorf("failed to connect metadata grpc service: %v", err)
		}
	}, nil
}
