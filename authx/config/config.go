package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Dsn string
}

func New(
	configName string,
	configType string,
	configPaths []string,
) (Config, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	for _, cp := range configPaths {
		viper.AddConfigPath(cp)
	}
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, nil
	}

	datasource := viper.GetStringMapString("app.datasource")
	var dsn = ""
	for key, value := range datasource {
		dsn += string(key + "=" + value + " ")
	}
	dsn = strings.TrimSpace(dsn)

	return Config{
		Dsn: dsn,
	}, nil
}
