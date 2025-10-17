package server

import (
	"fiber-jwt-demo/internal/auth"
	"fiber-jwt-demo/internal/logger"
	"fiber-jwt-demo/internal/race"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	app.Use(logger.LoggerMiddleware())

	app.Post("/login", auth.LoginHandler)

	api := app.Group("/api", auth.AuthMiddleware)
	api.Get("/profile", auth.ProfileHandler)
	api.Get("/race", race.RaceHandler)

	return app
}
