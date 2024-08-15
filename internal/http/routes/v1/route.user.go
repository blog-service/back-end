package v1

import (
	handler "back-end/internal/http/handlers/v1"
	"back-end/internal/http/middlewares"
	"github.com/gin-gonic/gin"
)

func userRoutes(g *gin.RouterGroup) {
	userGroups := g.Group("/users")
	authRoutes(userGroups)
}

func authRoutes(g *gin.RouterGroup) {
	userHandler := handler.NewUserHandler()
	g.POST("/sign-up", userHandler.SignUp)
	g.POST("/sign-in", userHandler.SignIn)

	g.Use(middlewares.ValidateAccess())
	g.GET("/info", userHandler.GetInfo)
}
