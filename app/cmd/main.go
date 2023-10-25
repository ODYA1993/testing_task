package main

import (
	"flag"
	"github.com/DmitryOdintsov/rest-api/app/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yaml", "path to config file")
}

func main() {
	config := apiserver.NewConfig()

	config.LoadConfigENV()
	//config.LoadConfigYAML(configPath)

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
