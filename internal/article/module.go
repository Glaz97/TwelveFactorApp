package article

import (
	"github.com/Glaz97/twelvefactorapp/internal/database"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("article",
	fx.Provide(
		func(lc fx.Lifecycle, db *database.Database, log *zap.Logger) (*Service, error) {
			s := NewArticleService(db, log)
			lc.Append(fx.Hook{OnStart: s.Start})
			return s, nil
		}),
)
