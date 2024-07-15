package main

import (
	"back-end/internal/config"
	"back-end/internal/datasource/database"
	"back-end/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	if err := config.InitConfig(); err != nil {
		logger.ConsoleLog().Error().Err(err).Msg("Init Config failed")
	}

	db, err := database.ConnectToDB(config.Cfg.MongodbUrl, config.Cfg.DatabaseName)
	if err != nil {
		logger.ConsoleLog().WithLevel(zerolog.FatalLevel).Err(err).Msg("Connect to mongodb failed")
	}

	fmt.Printf("Connected to MongoDB at %s\n", db.Name())

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	if err = r.Run(); err != nil {
		logger.ConsoleLog().Error().Err(err).Msg("Run err")
	}
}
