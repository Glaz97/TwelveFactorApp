package v1

import (
	"go.uber.org/fx"
)

var Module = fx.Module("v1",
	fx.Provide(
		NewArticleHandler,
		NewRouter,
	),
)
