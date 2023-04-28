package config

import (
	"strings"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Db        *gorm.DB
	TokenAuth *jwtauth.JWTAuth
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
		return Config{}, err
	}

	var dsn = ""
	for key, value := range viper.GetStringMapString("app.datasource") {
		dsn += string(key + "=" + value + " ")
	}
	db, err := gorm.Open(
		postgres.Open(strings.TrimSpace(dsn)),
		&gorm.Config{},
	)
	if err != nil {
		return Config{}, err
	}

	jwta := jwtauth.New(
		"HS256",
		[]byte(viper.GetString("app.security.jwt-secret")),
		nil,
	)

	return Config{
		Db:        db,
		TokenAuth: jwta,
	}, nil
}
