package loader

import (
	"log"
	"os"

	"github.com/abdheshnayak/ur-proxy/entity"
	g "github.com/abdheshnayak/ur-proxy/global"
	"sigs.k8s.io/yaml"
)

func GetConfiguration() (*entity.RoutesConfig, error) {

	var config entity.RoutesConfig

	s := os.Getenv("CONFIG_FILE")
	if s == "" {
		s = "./routes.yml"
	}

	b, err := os.ReadFile(s)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil

}

func loadConfigurations() error {
	config, err := GetConfiguration()
	if err != nil {
		return err
	}
	g.SetConfig(config)

	return nil
}

func StartLoading() {
	go func() {
		// for {
		// 	// log.Println("loaded")
		// 	time.Sleep(2 * time.Second)
		// }

		if err := loadConfigurations(); err != nil {
			log.Println(err)
		}

		log.Println("loaded")
	}()
}
