package routes

import (
	"dhikrama.com/web/src/handler/user"
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(c *fiber.App) {
	api := c.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", user.UserHandlerRead)
	v1.Get("/user/:id", user.UserGetById)
	v1.Post("/user", user.UserHandlerCreate)
	v1.Put("/user/:id", user.UserHandlerUpdate)
	v1.Delete("/user/:id", user.UserHandlerDeleteId)

}
