package main

import (
	"github.com/abdheshnayak/ur-proxy/handler"
	"github.com/abdheshnayak/ur-proxy/loader"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.All("*", handler.HandleRequest)

	loader.StartLoading()

	app.Listen(":4000")
}
