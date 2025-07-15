package handler

import (
	v1 "github.com/Glaz97/twelvefactorapp/internal/handler/v1"
	"go.uber.org/fx"
)

var Module = fx.Module("handler",
	v1.Module,
	fx.Provide(
		NewHandler,
	),
	fx.Invoke(
		SetProductionMode,
	),
)
