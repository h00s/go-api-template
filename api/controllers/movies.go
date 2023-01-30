package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s/go-api-template/api/models"
	"github.com/h00s/go-api-template/api/services"
)

func GetMoviesHandler(c *fiber.Ctx) error {
	m := models.NewMovies(services.GetServices(c))
	movies, err := m.FindAll()
	if err != nil {
		services.GetServices(c).Logger.Println(err)
		return c.SendStatus(500)
	}
	return c.JSON(movies)
}

func CreateMovieHandler(c *fiber.Ctx) error {
	s := services.GetServices(c)
	m := models.NewMovies(s)
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
