package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/kwstars/film-hive/api/rating/service/v1"
	"github.com/kwstars/film-hive/app/rating/service/internal/conf"
	"github.com/kwstars/film-hive/app/rating/service/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, rating *service.RatingService, _ log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			// logging.Server(logger),
			ratelimit.Server(),
			tracing.Server(),
		),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRatingServiceHTTPServer(srv, rating)
	return srv
}
