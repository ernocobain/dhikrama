package htmlhandler

import "github.com/gofiber/fiber/v2"

func IndexHandler(ctx *fiber.Ctx) error {
	// ctx.Request().Header.Set(fiber.HeaderAcceptEncoding, "br")

	return ctx.Render("index/index", fiber.Map{
		"title": "Welcome to my website",
	}, "layouts/main")
}
