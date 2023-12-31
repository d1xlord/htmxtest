package routes

import (
	"htmxtest/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	handlers := handlers.NewFilmsHandler()
	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.HandleFilms(c)
	})
	app.Post("/film", func(c *fiber.Ctx) error {
		return handlers.HandleAddFilm(c)
	})
}
