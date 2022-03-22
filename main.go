package main

import (
	"net/http"

	"github.com/gofiber/template/html"

	"dhikrama.com/web/src/model/database"
	"dhikrama.com/web/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	database.DatabaseMysql()
	database.RunMigration()
	engine := html.NewFileSystem(http.Dir("./public/views"), ".html")
	app := fiber.New(
		fiber.Config{
			Views:   engine,
			Prefork: false,
		},
	)
	app.Use(compress.New())
	routes.RoutesInit(app)
	routes.RoutesStaticInit(app)

	app.Listen(":8080")
}
