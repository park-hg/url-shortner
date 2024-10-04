package internal

import (
	"github.com/gofiber/fiber/v2"

	"traffic-reporter/internal/shortener/controller"
)

func RegisterRouters(router *fiber.App, app *App) {
	shortenURlController := controller.NewShortenURLController(app.config.ServerConfig, app.ShortenURLUseCase)
	router.Post("/admin/urls/shorten", shortenURlController.GenerateShortURL)
	router.Get("/admin/urls/shortened", shortenURlController.GetShortened)
	router.Get("/:shortened", shortenURlController.RedirectToOriginal)
}
