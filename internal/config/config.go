package config

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

func DefaultConfig() Config {
	return Config{
		MongoDB: MongoDB{
			URI:             "mongodb://localhost:27017",
			Database:        "twelvefactorapp",
			ReadPrefNearest: true,
		},
		HTTP: HTTP{
			Address:      "0.0.0.0:8056",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			CORS: CORS{
				AllowOrigins: []string{
					"*",
				},
				AllowHeaders: []string{"*"},
			},
		},
		Logger: Logger{Debug: false},
		Swagger: Swagger{
			BasePath: "/",
		},
	}
}

type Config struct {
	HTTP    HTTP    `mapstructure:"http"`
	Logger  Logger  `mapstructure:"logger"`
	Swagger Swagger `mapstructure:"swagger"`
	MongoDB MongoDB `mapstructure:"mongodb"`
}

// NewConfig reads configuration from YAML files and environment variables with provided defaults.
// It sets the values in following order:
//   - Values from `defaults`.
//   - Values from yaml file.
//   - Values from environment variables, for example DB_PORT=5000 for `db.port` key.
func NewConfig() (*Config, error) {
	v, err := load(DefaultConfig())
	if err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to umarshal config, %w", err)
	}
	return &config, nil
}

// YAML returns the config as YAML.
func (c *Config) YAML() ([]byte, error) {
	m := make(map[string]any)
	if err := mapstructure.Decode(c, &m); err != nil {
		return nil, err
	}
	return yaml.Marshal(m)
}

type MongoDB struct {
	URI      SecretString `mapstructure:"uri"`
	Database string       `mapstructure:"database"`

	// ReadPrefNearest enables the nearest read preference in MongoDB.
	// This is useful for read-heavy deployments distributed across multiple regions.
	ReadPrefNearest bool `mapstructure:"read_pref_nearest"`
}

type HTTP struct {
	Address      string        `mapstructure:"address"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	CORS         CORS          `mapstructure:"cors"`
}

type CORS struct {
	AllowOrigins []string `mapstructure:"allow_origins"`
	AllowHeaders []string `mapstructure:"allow_headers"`
}

type Logger struct {
	Debug bool `mapstructure:"debug"`
}

type Swagger struct {
	BasePath string `mapstructure:"base_path"`
}

type SecretString string

func (s SecretString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Mask())
}

func (s SecretString) Mask() string {
	if len(s) <= 6 {
		return string(s)
	}

	return string(s[:3]) + "*****" + string(s[len(s)-3:])
}
