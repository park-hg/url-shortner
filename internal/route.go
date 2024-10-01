package internal

import (
	"github.com/gofiber/fiber/v2"

	"traffic-reporter/internal/shortener/controller"
)

func RegisteRouters(router *fiber.App, app *App) {
	shortenURlController := controller.NewShortenURLController(app.ShortenURLUseCase)

	router.Post("/shorten", shortenURlController.Shorten)
}
