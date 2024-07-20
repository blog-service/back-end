package config

import (
	"back-end/pkg/logger"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var consoleLog = logger.ConsoleLog()

type Configuration struct {
	Port           int    `env:"PORT" envDefault:"8080"`
	MongodbUrl     string `env:"MONGODB_URL" envDefault:"mongodb://localhost:27017/"`
	DatabaseName   string `env:"DATABASE_NAME" envDefault:"localdb"`
	PrivateKeyPath string `env:"PRIVATE_KEY_PATH" envDefault:"certs/privateKey.pem"`
	PublicKeyPath  string `env:"PUBLIC_KEY_PATH" envDefault:"certs/publicKey.pem"`
}

var config *Configuration

func GetConfig() Configuration {
	if config == nil {
		if err := godotenv.Load(); err != nil {
			consoleLog.Fatal().Err(err).Str("func", "GetConfig-godotenv.Load").Msg("config")
		}
		config = &Configuration{}
		if err := env.Parse(config); err != nil {
			consoleLog.Fatal().Err(err).Str("func", "GetConfig-envParse").Msg("config")
		}
	}
	return *config
}
