package main

import (
	"htmxtest/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
