package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"

	"traffic-reporter/config"
	"traffic-reporter/internal"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	done := make(chan struct{})

	c := config.MustNewConfig()

	app := internal.InitApp(c)
	srv := fiber.New(c.FiberConfig)
	internal.RegisterRouters(srv, app)

	go func() {
		_ = <-sigChan
		fmt.Println("gracefully shutting down...")
		_ = srv.Shutdown()

		done <- struct{}{}
	}()

	if err := srv.Listen(":8080"); err != nil {
		panic(err)
	}

	<-done
	_ = app.Teardown()
}
