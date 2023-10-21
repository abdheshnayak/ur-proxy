package handler

import (
	"log"

	g "github.com/abdheshnayak/ur-proxy/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func HandleRequest(c *fiber.Ctx) error {
	g.GCtx.Mu.RLock()
	defer g.GCtx.Mu.RUnlock()

	hostPath := c.Hostname() + c.Path()
	config, ok := g.GCtx.Routes[hostPath]
	if ok && config.Active {

		log.Println(c.GetReqHeaders()["Authorization"])

		f := proxy.Forward(config.Backend.String())
		return f(c)

	}
	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}
