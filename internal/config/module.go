package config

import (
	"encoding/json"
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("config",
	fx.Provide(
		NewConfig,
		Destructure,
	),
	fx.Invoke(
		LogConfig,
	),
)

func Destructure(cfg *Config) (
	*MongoDB,
	*HTTP,
	*Logger,
	*Swagger,
	*CORS,
) {
	return &cfg.MongoDB,
		&cfg.HTTP,
		&cfg.Logger,
		&cfg.Swagger,
		&cfg.HTTP.CORS
}

func LogConfig(cfg *Config, log *zap.Logger) error {
	j, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	log.Info("Configuration loaded", zap.String("config", string(j)))
	return nil
}
