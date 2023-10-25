package apiserver

import (
	"flag"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	BinAddr     string `env:"bin_addr" yaml:"bin_addr"`
	LogLevel    string `env:"log_level" yaml:"log_level"`
	DatabaseURL string `env:"database_url" yaml:"database_url"`
	SessionKey  string `env:"session_key" yaml:"session_key"`
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
	}
}

func (c *Config) LoadConfigENV() {
	err := godotenv.Load("configs/app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.BinAddr = os.Getenv("bin_addr")
	c.LogLevel = os.Getenv("log_level")
	c.SessionKey = os.Getenv("session_key")
	c.DatabaseURL = os.Getenv("database_url")
}

func (c *Config) LoadConfigYAML(configPath string) {
	flag.Parse()
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Println(err)
	}
}
