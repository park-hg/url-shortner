package main

import (
	"github.com/gofiber/fiber/v2"

	"traffic-reporter/config"
	"traffic-reporter/internal"
)

func main() {
	c := config.MustNewConfig()

	app := internal.InitApp(c)
	srv := fiber.New(c.FiberConfig)
	internal.RegisterRouters(srv, app)
	if err := srv.Listen(":8080"); err != nil {
		panic(err)
	}
}
