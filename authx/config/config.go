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

	db, err := CreateDb("app.datasource")
	if err != nil {
		return Config{}, err
	}

	jwta := CreateTokenAuth("app.security.jwt-secret")

	return Config{
		Db:        db,
		TokenAuth: jwta,
	}, nil
}

func CreateDb(configKey string) (*gorm.DB, error) {
	var dsn = ""
	for key, value := range viper.GetStringMapString(configKey) {
		dsn += string(key + "=" + value + " ")
	}
	return gorm.Open(
		postgres.Open(strings.TrimSpace(dsn)),
		&gorm.Config{},
	)
}

func CreateTokenAuth(configKey string) *jwtauth.JWTAuth {
	return jwtauth.New(
		"HS256",
		[]byte(viper.GetString(configKey)),
		nil,
	)
}
