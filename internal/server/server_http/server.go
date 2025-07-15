package server_http

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/Glaz97/twelvefactorapp/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	*http.Server
	log *zap.Logger
}

func NewServer(
	cfg *config.HTTP,
	handler http.Handler,
	log *zap.Logger,
) *Server {
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
	return &Server{
		Server: srv,
		log:    log,
	}
}

func (s *Server) Start(_ context.Context) error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	s.log.Info("Starting HTTP server", zap.String("addr", s.Addr))
	go func() {
		if err := s.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Error("HTTP server error", zap.Error(err))
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("Shutting down http server")
	return s.Shutdown(ctx)
}
