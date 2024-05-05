package usersController

func Init(g *fiber.Group) {
	var (
		h = userhandler.New()
	)
	publicRoute := g.Group("/users")
	privateRoute := publicRoute
	privateRoute.POST("/register", h.Register)
}
