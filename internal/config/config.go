package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type config struct {
	Port         int    `env:"PORT" envDefault:"8080"`
	MongodbUrl   string `env:"MONGODB_URL" envDefault:"mongodb://localhost:27017/"`
	DatabaseName string `env:"DATABASE_NAME" envDefault:"localdb"`
}

var Cfg = &config{}

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	if err := env.Parse(Cfg); err != nil {
		return err
	}
	return nil
}
