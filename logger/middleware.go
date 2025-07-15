package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// StructuredLogger is a middleware that logs requests using zap.
// Inspired by https://learninggolang.com/it5-gin-structured-logging.html
func StructuredLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC() // Start timer

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next() // Process request

		latency := time.Since(start) // Stop timer

		args := []zap.Field{
			zap.String("client_ip", c.RemoteIP()),
			zap.String("method", c.Request.Method),
			zap.Int("status_code", c.Writer.Status()),
			zap.Int("body_size", c.Writer.Size()),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("latency", latency.String()),
			zap.Int64("latency_ms", latency.Milliseconds()),
			zap.String("user_agent", c.Request.UserAgent()),
		}

		msg := c.Errors.ByType(gin.ErrorTypePrivate).String()
		if c.Writer.Status() >= 500 {
			logger.Error(msg, args...)
		} else {
			logger.Info(msg, args...)
		}
	}
}
