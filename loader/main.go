package loader

import (
	"log"
	"os"
	"time"

	"github.com/abdheshnayak/ur-proxy/entity"
	g "github.com/abdheshnayak/ur-proxy/global"
	"sigs.k8s.io/yaml"
)

func getConfiguration() (*entity.RoutesConfig, error) {

	var config entity.RoutesConfig

	b, err := os.ReadFile("./routes.yml")
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil

}

func loadConfigurations() error {
	g.GCtx.Mu.Lock()
	defer g.GCtx.Mu.Unlock()

	config, err := getConfiguration()
	if err != nil {
		return err
	}
	g.SetConfig(config)

	return nil
}

func StartLoading() {
	go func() {
		for {
			if err := loadConfigurations(); err != nil {
				log.Println(err)
			}
			// log.Println("loaded")
			time.Sleep(2 * time.Second)
		}
	}()
}
