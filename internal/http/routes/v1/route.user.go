package v1

func (r *router) userRoutes() {
	r.engine.Group("/users")
}
