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
