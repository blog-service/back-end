package v1

import handler "back-end/internal/http/handlers/v1"

func (r *router) userRoutes() {
	r.engine.Group("/users")
}

func (r *router) authRoutes() {
	userHandler := handler.NewUserHandler()
	r.engine.POST("/sign-up", userHandler.SignUp)
}
