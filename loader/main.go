package loader

import (
	"net/url"
	"time"

	g "github.com/abdheshnayak/ur-proxy/global"
)

func loadConfigurations() {
	// Load, add, or remove configurations as necessary
	g.GCtx.Mu.Lock()
	defer g.GCtx.Mu.Unlock()

	// Example of adding a route
	backendURL, _ := url.Parse("http://localhost:3000")
	g.GCtx.Routes["archlinux:4000/"] = &g.RouteConfig{
		Active:  true,
		Backend: backendURL,
	}
}

func StartLoading() {
	go func() {
		for {
			loadConfigurations()
			time.Sleep(30 * time.Second)
		}
	}()
}
