package server

import (
	v1 "github.com/kwstars/film-hive/api/metadata/service/v1"
	"github.com/kwstars/film-hive/app/metadata/service/internal/conf"
	"github.com/kwstars/film-hive/app/metadata/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, metadata *service.MetadataService, _ log.Logger) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			// logging.Server(logger),
			ratelimit.Server(),
			tracing.Server(),
		),
	}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	// v1.RegisterGreeterServer(srv, metadata)
	v1.RegisterMetadataServiceServer(srv, metadata)
	return srv
}
