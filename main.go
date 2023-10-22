package main

import (
	"github.com/abdheshnayak/ur-proxy/handler"
	"github.com/abdheshnayak/ur-proxy/loader"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(
		logger.New(
			logger.Config{
				Format:     "${time} ${status} - ${method} ${latency} \t ${path} \n",
				TimeFormat: "02-Jan-2006 15:04:05",
				TimeZone:   "Asia/Kolkata",
			},
		),
	)

	app.All("*", handler.HandleRequest)

	loader.StartLoading()

	if err := app.Listen(":4000"); err != nil {
		panic(err)
	}
}
