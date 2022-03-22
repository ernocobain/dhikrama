package htmlhandler

import "github.com/gofiber/fiber/v2"

func ContactHandler(ctx *fiber.Ctx) error {
	// ctx.Request().Header.Set(fiber.HeaderAcceptEncoding, "br")
	return ctx.Render("about/about", fiber.Map{
		"title": "Welcome to my Contact",
	}, "layouts/main")
}
