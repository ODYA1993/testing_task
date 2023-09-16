package main

import (
	"flag"
	"github.com/DmitryOdintsov/rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config file")
}

func main() {
	config := apiserver.NewConfig()
	config.LoadConfigENV()
	//config.LoadConfigYAML(configPath)

	server := apiserver.NewAPIServer(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
