package v1

import "github.com/gin-gonic/gin"

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
	r.userRoutes()
}
