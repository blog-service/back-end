package v1

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	V1()
}

type router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) Router {
	return &router{
		engine: engine,
	}
}

func (r *router) V1() {
	v1 := r.engine.Group("/api/blog-service/v1/")
	userRoutes(v1)
}
