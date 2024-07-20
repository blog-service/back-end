package main

import (
	"back-end/internal/datasource/database"
	routes "back-end/internal/http/routes/v1"
	"back-end/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.ConnectToDB(); err != nil {
		logger.ConsoleLog().Fatal().Err(err)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	routes.NewRouter(r).V1()

	if err := r.Run(); err != nil {
		logger.ConsoleLog().Error().Str("func", "main-r.Run").Err(err).Msg("Run err")
	}
}
