package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"github.com/lftzzzzfeng/fasms/db/pg"
	httpserver "github.com/lftzzzzfeng/fasms/server"
)

// Config is the main application configurations.
type Config struct {
	EnvName  string
	Database *pg.PGConnectionConfig `yaml:"database"`
	Server   *httpserver.Config     `yaml:"server"`
}

func Load(env string) (*Config, error) {
	viper.AddConfigPath(fmt.Sprintf("./conf"))
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	errViper := viper.ReadInConfig()
	if errViper != nil {
		return nil, fmt.Errorf("fatal error config file: %w", errViper)
	}

	var c Config

	errM := viper.Unmarshal(&c)
	if errM != nil {
		return nil, fmt.Errorf("unable to decode into struct, %w", errM)
	}

	return &c, nil
}
