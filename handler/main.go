package handler

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/abdheshnayak/ur-proxy/loader"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func checkAuth(url string, method string, path string, header *fasthttp.RequestHeader) (*fasthttp.Response, error) {
	url = strings.ReplaceAll(url, "{method}", method)
	url = strings.ReplaceAll(url, "{path}", path)

	// Creating a new fasthttp request and response
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	// defer fasthttp.ReleaseResponse(resp)

	// Setting the request method and URL
	req.Header.SetMethod("GET")
	req.SetRequestURI(url)

	header.VisitAll(func(key, value []byte) {
		req.Header.SetBytesKV(key, value)
	})

	err := fasthttp.Do(req, resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}

func HandleRequest(c *fiber.Ctx) error {

	// g.GCtx.Mu.RLock()
	// defer g.GCtx.Mu.RUnlock()

	randId := uuid.New().String()
	log.Println("processing started:", randId)
	defer log.Println("processing ended:", randId)

	hostname := c.Hostname()
	path := c.Path()
	method := string(c.Request().Header.Method())
	header := &c.Request().Header

	rc, err := loader.GetConfiguration()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	for _, rc := range rc.Routes {
		if rc.Host == hostname {
			for _, hp := range rc.Paths {
				re := regexp.MustCompile(hp.Path)
				matched := re.MatchString(path)

				if matched {
					// addr := fmt.Sprintf("http://%s:%d%s", hp.Backend.Service.Name, hp.Backend.Service.Port, path)

					if rc.AuthUrl != nil {

						resp, err := checkAuth(*rc.AuthUrl, method, path, header)
						if err != nil {
							return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
						}

						if resp.StatusCode() != fiber.StatusOK {
							defer fasthttp.ReleaseResponse(resp)
							log.Println("auth failed", resp.StatusCode())
							resp.Header.VisitAll(func(key, value []byte) {
								c.Response().Header.Add(string(key), string(value))
							})
							return c.Status(resp.StatusCode()).SendString(string(resp.Body()))
						}
					}

					// return proxy.Do(c, addr)

					f := proxy.Balancer(proxy.Config{
						Servers: []string{
							fmt.Sprintf("http://%s:%d", hp.Backend.Service.Name, hp.Backend.Service.Port),
						},
					})

					// f := proxy.Forward(addr)
					return f(c)
				}
			}

			log.Println("no matches found for", hostname, path)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}
