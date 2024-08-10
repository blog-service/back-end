package main

import (
	"fmt"

	"back-end/internal/config"
	"back-end/internal/datasource/database"
	"back-end/internal/http/middlewares"
	routes "back-end/internal/http/routes/v1"
	"back-end/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.ConnectToDB(); err != nil {
		logger.ConsoleLog().Fatal().Err(err)
	}

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.Use(gin.Recovery())

	routes.NewRouter(r).V1()

	if err := r.Run(fmt.Sprintf(":%d", config.GetConfig().Port)); err != nil {
		logger.ConsoleLog().Error().Str("func", "main-r.Run").Err(err).Msg("Run err")
	}
}
