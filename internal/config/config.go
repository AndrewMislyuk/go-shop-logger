package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB     DB
	Server Server
}

type DB struct {
	URI      string `envconfig:"DB_URI"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
	Database string `envconfig:"DB_DATABASE"`
}

type Server struct {
	Port int `envconfig:"SERVER_PORT"`
}

func New() (cfg Config, err error) {
	viper.SetConfigFile("./.env")
	err = viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	cfg.Server.Port = viper.GetInt("SERVER_PORT")

	cfg.DB.URI = viper.GetString("DB_URI")
	cfg.DB.Username = viper.GetString("DB_USERNAME")
	cfg.DB.Password = viper.GetString("DB_PASSWORD")
	cfg.DB.Database = viper.GetString("DB_DATABASE")

	return cfg, nil
}
