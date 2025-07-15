package app

import (
	"github.com/Glaz97/twelvefactorapp/internal/article"
	"github.com/Glaz97/twelvefactorapp/internal/config"
	"github.com/Glaz97/twelvefactorapp/internal/database"
	"github.com/Glaz97/twelvefactorapp/internal/handler"
	"github.com/Glaz97/twelvefactorapp/internal/server/server_http"
	"github.com/Glaz97/twelvefactorapp/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.WithLogger(logger.FxLogger),
	config.Module,
	logger.Module,
	database.Module,
	handler.Module,
	server_http.Module,
	article.Module,
)
