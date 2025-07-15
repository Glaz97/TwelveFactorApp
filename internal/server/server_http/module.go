package server_http

import (
	"net/http"

	"github.com/Glaz97/twelvefactorapp/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("server_http",
	fx.Provide(newServer),
)

func newServer(
	lc fx.Lifecycle,
	cfg *config.HTTP,
	handler http.Handler,
	log *zap.Logger,
) *Server {
	srv := NewServer(cfg, handler, log)
	lc.Append(fx.Hook{
		OnStart: srv.Start,
		OnStop:  srv.Stop,
	})
	return srv
}
