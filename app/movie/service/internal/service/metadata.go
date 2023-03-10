package service

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	metadata "github.com/kwstars/film-hive/api/metadata/service/v1"
	"github.com/kwstars/film-hive/app/movie/service/internal/conf"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"time"
)

func newMetadataClient(c *conf.Bootstrap, rc naming_client.INamingClient) metadata.MetadataServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("discovery:///metadata.grpc"),
		grpc.WithDiscovery(nacos.New(rc)),
	)
	if err != nil {
		panic(err)
	}
	return metadata.NewMetadataServiceClient(conn)
}
