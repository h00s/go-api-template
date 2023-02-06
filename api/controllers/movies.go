package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s/go-api-template/api/middleware"
	"github.com/h00s/go-api-template/api/models"
)

func GetMoviesHandler(c *fiber.Ctx) error {
	m := middleware.GetModels(c).Movies
	movies, err := m.FindAll()
	if err != nil {
		middleware.GetServices(c).Logger.Println(err)
		return c.SendStatus(500)
	}
	return c.JSON(movies)
}

func CreateMovieHandler(c *fiber.Ctx) error {
	s := middleware.GetServices(c)
	m := middleware.GetModels(c).Movies
	movie := &models.Movie{}
	if err := c.BodyParser(movie); err != nil {
		s.Logger.Println(err)
		return c.SendStatus(400)
	}
	if err := m.Create(movie); err != nil {
		s.Logger.Println(err)
		return c.SendStatus(500)
	}
	return c.JSON(movie)
}
