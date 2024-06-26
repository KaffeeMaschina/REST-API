package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env           string `yaml:"env" env-default:"local" env-required:"true"`
	Storagecfg    `yaml:"storage"`
	HTTPServer    `yaml:"http_server"`
	NatsStreaming `yaml:"nats_streaming"`
}
type Storagecfg struct {
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
}
type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}
type NatsStreaming struct {
	ClusterID string `yaml:"cluster_id" env-required:"true"`
	CLientID  string `yaml:"client_id" ent-required:"true"`
}

func MustLoad() *Config {
	a := godotenv.Load()
	_ = a
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s doesn't exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %s", err)
	}
	return &cfg
}
