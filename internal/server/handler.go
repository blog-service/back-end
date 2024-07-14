package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	log.Fatal("Hello, welcome to blog services with gin and mongoDB", c.Request);
}