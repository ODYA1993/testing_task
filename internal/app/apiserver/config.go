package apiserver

import (
	"flag"
	"github.com/DmitryOdintsov/rest-api/internal/app/store"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	BinAddr  string          `env:"bin_addr" yaml:"bin_addr"`
	LogLevel string          `env:"log_level" yaml:"log_level"`
	Store    *store.ConfigDB `env:"store" yaml:"store"`
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfigDB(),
	}
}

func (c *Config) LoadConfigENV() {

	err := godotenv.Load("configs/app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.BinAddr = os.Getenv("bin_addr")
	c.LogLevel = os.Getenv("log_level")
	c.Store.Host = os.Getenv("host")
	c.Store.DBName = os.Getenv("db_name")
	c.Store.Port = os.Getenv("port")
	c.Store.UserName = os.Getenv("user_name")
	c.Store.Password = os.Getenv("password")
	c.Store.SslMode = os.Getenv("sslmode")
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
