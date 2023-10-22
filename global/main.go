package global

import (
	"sync"

	"github.com/abdheshnayak/ur-proxy/entity"
)

type GContext struct {
	Config entity.RoutesConfig
	Mu     sync.RWMutex
}

var (
	GCtx = GetGContext()
)

func SetConfig(config *entity.RoutesConfig) {
	GCtx.Config = *config
}

func GetGContext() *GContext {
	return &GContext{
		Config: entity.RoutesConfig{},
		Mu:     sync.RWMutex{},
	}
}
