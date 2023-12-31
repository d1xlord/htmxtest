package handlers

import (
	"fmt"
	"htmxtest/src/models"
	"htmxtest/src/templates"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/utils"
)

type filmsHandler struct {
	films []*models.Film
}

func NewFilmsHandler() filmsHandler {
	handler := filmsHandler{}
	handler.films = []*models.Film{
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
		{Title: "Schindler's List", Director: "Steven Spielberg"},
	}
	return handler
}

func (f *filmsHandler) HandleFilms(c *fiber.Ctx) error {
	newList := f.getFilmsList()
	return c.Render("films", fiber.Map{
		"Films": newList,
	})
}

func (f *filmsHandler) HandleAddFilm(c *fiber.Ctx) error {
	title := utils.CopyString(c.FormValue("title"))
	director := utils.CopyString(c.FormValue("director"))
	if title != "" && director != "" {
		f.addToFilmsList(title, director)
	}
	handler := adaptor.HTTPHandler(
		templ.Handler(templates.Films(f.getFilmsList())),
	)
	return handler(c)
}

func (f *filmsHandler) getFilmsList() []*models.Film {
	return f.films
}

func (f *filmsHandler) addToFilmsList(title string, director string) error {
	f.films = append(f.films, &models.Film{Title: title, Director: director})
	return nil
}

func (f *filmsHandler) printFilmsList() {
	for _, film := range f.films {
		fmt.Println(film.Title, film.Director)
	}
}
