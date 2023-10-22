package entity

type RoutesConfig struct {
	Version string        `json:"version"`
	Routes  []RouteConfig `json:"routes"`
}

type RouteConfig struct {
	Host    string  `json:"host"`
	Paths   []Path  `json:"paths"`
	AuthUrl *string `json:"authUrl"`
}

type Path struct {
	Path     string `json:"path"`
	PathType string `json:"pathType"`
	Backend  struct {
		Service struct {
			Name string `json:"name"`
			Port int64  `json:"port"`
		} `json:"service"`
	} `json:"backend"`
}
