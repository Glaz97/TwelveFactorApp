package database

import (
	"context"

	"github.com/Glaz97/twelvefactorapp/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("database",
	fx.Provide(
		newDatabase,
	),
)

func newDatabase(lc fx.Lifecycle, cfg *config.MongoDB, log *zap.Logger) (*Database, error) {
	db, err := NewDatabase(cfg, log)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Connecting to database")
			return db.Start(ctx)
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Disconnecting from database")
			return db.Stop(ctx)
		},
	})

	return db, nil
}
