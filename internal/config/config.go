package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port           string `yaml:"port"`
	DatabaseConfig `yaml:"database_config"`
}
type DatabaseConfig struct {
	Host     string `yaml:"host" env-default:"localhost"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"potgres"`
	DBName   string `yaml:"dbname" env-default:"lexoread"`
	DBPort   string `yaml:"port" env-default:"5432"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

func MustLoad(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	return &cfg
}
