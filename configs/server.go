package configs

import (
	"taufikRahadi/fiber-boilerplate/application"

	"github.com/gofiber/fiber/v2"
)

func InitApi() *fiber.App {
	api := fiber.New()

	apiGroup := api.Group("/api")

	// INIT Application
	application.ApplicationModule(apiGroup)

	return api
}
