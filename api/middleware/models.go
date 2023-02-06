package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s/go-api-template/api/models"
	"github.com/h00s/go-api-template/services"
)

type ModelsMiddleware struct {
	Movies *models.Movies
}

func NewModelsMiddleware(services *services.Services) *ModelsMiddleware {
	return &ModelsMiddleware{
		Movies: models.NewMovies(services),
	}
}

func (m *ModelsMiddleware) ModelsMiddleware(c *fiber.Ctx) error {
	c.Locals("models", m)
	return c.Next()
}

func GetModels(c *fiber.Ctx) *ModelsMiddleware {
	return c.Locals("models").(*ModelsMiddleware)
}
