package logger

import (
	"context"
	"errors"
	"syscall"

	"github.com/Glaz97/twelvefactorapp/internal/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("logger",
	fx.Provide(Logger),
)

func FxLogger(log *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log}
}

func Logger(lc fx.Lifecycle, cfg *config.Logger) (*zap.Logger, error) {
	var log *zap.Logger
	var err error

	if cfg.Debug {
		log, err = zap.NewDevelopment()
	} else {
		log, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			err := log.Sync()
			// See the thread: https://github.com/uber-go/zap/issues/991#issuecomment-962098428
			if errors.Is(err, syscall.ENOTTY) {
				return nil
			}

			return err
		},
	})

	return log, nil
}
