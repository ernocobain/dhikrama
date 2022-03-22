package routes

import (
	htmlhandler "dhikrama.com/web/src/html_handler"
	"github.com/gofiber/fiber/v2"
)

func RoutesStaticInit(r *fiber.App) {
	r.Static("/", "./public/static", fiber.Static{
		Compress: false,
	})
	r.Get("/", htmlhandler.IndexHandler)
	r.Get("/about", htmlhandler.AboutHandler)
	r.Get("/contact", htmlhandler.ContactHandler)
}
