package global

import (
	"net/url"
	"sync"
)

type RouteConfig struct {
	Active  bool
	Backend *url.URL
}

type GContext struct {
	Routes map[string]*RouteConfig
	Mu     sync.RWMutex
}

var (
	GCtx = GetGContext()
)

func GetGContext() *GContext {
	return &GContext{
		Routes: make(map[string]*RouteConfig),
		Mu:     sync.RWMutex{},
	}
}
