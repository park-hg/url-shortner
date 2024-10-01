package controller

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"traffic-reporter/config"
	"traffic-reporter/internal/pkg"
	"traffic-reporter/internal/shortener/usecase"
)

type ShortenURLController struct {
	host string
	uc   *usecase.ShortenURLUseCase
}

func NewShortenURLController(srvConfig config.ServerConfig, uc *usecase.ShortenURLUseCase) *ShortenURLController {
	return &ShortenURLController{host: srvConfig.Host, uc: uc}
}

type ShortenURLRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenURLResponse struct {
	ShortURL string `json:"short_url"`
}

func (c *ShortenURLController) Shorten(fc *fiber.Ctx) error {
	ctx := fc.UserContext()
	req := ShortenURLRequest{}
	if err := fc.BodyParser(&req); err != nil {
		return fc.Status(http.StatusBadRequest).JSON(pkg.FiberError{
			Error:       err,
			Description: "invalid request body",
		})
	}
	if req.OriginalURL == "" {
		return fc.Status(http.StatusBadRequest).JSON(pkg.FiberError{
			Error:       errors.New("original_url is empty"),
			Description: "invalid request body",
		})
	}

	shortened, err := c.uc.Shorten(ctx, req.OriginalURL)
	if err != nil {
		return fc.Status(http.StatusInternalServerError).JSON(pkg.FiberError{
			Error:       err,
			Description: "failed to create shortened url",
		})
	}

	return fc.Status(http.StatusCreated).JSON(ShortenURLResponse{ShortURL: c.host + "/" + shortened})
}

func (c *ShortenURLController) RedirectToOriginal(fc *fiber.Ctx) error {
	ctx := fc.UserContext()
	shortened := fc.Params("shortened")
	originalURL, err := c.uc.RetrieveOriginal(ctx, shortened)
	if err != nil {
		return fc.Status(http.StatusInternalServerError).JSON(pkg.FiberError{
			Error:       err,
			Description: "failed to retrieve original url",
		})
	}

	// TODO: switch status code via user's permission
	return fc.Redirect(originalURL, http.StatusFound)
}
