package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s/go-api-template/api/services"
)

type Models struct {
	Movies *Movies
}

func NewModels(services *services.Services) *Models {
	return &Models{
		Movies: NewMovies(services),
	}
}

func (m *Models) ModelsMiddleware(c *fiber.Ctx) error {
	c.Locals("models", m)
	return c.Next()
}

func GetModels(c *fiber.Ctx) *Models {
	return c.Locals("models").(*Models)
}
