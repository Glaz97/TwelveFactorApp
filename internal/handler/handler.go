package handler

import (
	"net/http"
	"time"

	"github.com/Glaz97/twelvefactorapp/internal/config"
	v1 "github.com/Glaz97/twelvefactorapp/internal/handler/v1"
	"github.com/Glaz97/twelvefactorapp/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	@title		TwelveFactorApp API
//	@version	1.0

func NewHandler(
	cfg *config.HTTP,
	log *zap.Logger,
	v1Handler *v1.Router,
	swagger *config.Swagger,
) (http.Handler, error) {
	router := gin.New()

	corsHandler, err := newCorsHandler(&cfg.CORS)
	if err != nil {
		return nil, err
	}

	// Middleware
	router.Use(
		gin.Recovery(),
		corsHandler,
		logger.StructuredLogger(log),
	)

	// System group
	{
		// Health check
		router.GET("/status", statusHandler)
	}

	// Stable API group
	v1Handler.RegisterRoutes(router)

	return router, nil
}

func SetProductionMode() {
	gin.SetMode(gin.ReleaseMode)
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func newCorsHandler(cfg *config.CORS) (gin.HandlerFunc, error) {
	corsCfg := cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowHeaders:     cfg.AllowHeaders,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowWildcard:    true,
		AllowMethods:     []string{"GET", "POST"},
	}
	err := corsCfg.Validate()
	if err != nil {
		return nil, err
	}
	return cors.New(corsCfg), err
}
