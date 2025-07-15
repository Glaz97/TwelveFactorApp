package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const (
	EnvPrefix         = "twelvefactorapp"
	DefaultConfigFile = "/etc/twelvefactorapp/twelvefactorapp.yaml"
	ConfigFileEnv     = "CONFIG_FILE"
	ConfigFileKey     = "config-file"
)

func load(defaults interface{}) (*viper.Viper, error) {
	v := viper.New()

	v.SetDefault(ConfigFileKey, DefaultConfigFile)
	err := v.BindEnv(ConfigFileKey, ConfigFileEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to bind config file env, %w", err)
	}
	v.SetConfigFile(v.GetString(ConfigFileKey))

	// NewConfig file if exists
	err = v.ReadInConfig()
	if err != nil {
		var notFoundErr viper.ConfigFileNotFoundError
		if !os.IsNotExist(err) && !errors.As(err, &notFoundErr) {
			return nil, fmt.Errorf("failed to read config, %w", err)
		}
	}

	v.AutomaticEnv()
	v.SetEnvPrefix(EnvPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	// set defaults
	defaultsMap := make(map[string]interface{})
	err = mapstructure.Decode(defaults, &defaultsMap)
	if err != nil {
		return nil, err
	}
	for key, value := range defaultsMap {
		v.SetDefault(key, value)
	}

	return v, nil
}
