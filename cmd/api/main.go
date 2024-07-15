package main

import (
	"back-end/internal/config"
	"back-end/internal/datasource/database"
	"back-end/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitConfig(); err != nil {
		logger.ConsoleLog().Error().Err(err).Str("func", "config.InitConfig").Msg("Init Config failed")
	}

	if err := database.ConnectToDB(); err != nil {
		logger.ConsoleLog().Fatal().Err(err)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	if err := r.Run(); err != nil {
		logger.ConsoleLog().Error().Str("func", "r.Run").Err(err).Msg("Run err")
	}
}
