package v1

func Init(e *fiber.App) {
	g := e.Group("/v1")
	userroute.Init(g)
}
