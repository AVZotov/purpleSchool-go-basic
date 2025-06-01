package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Key string
}

func NewEnvConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}
	key := os.Getenv("KEY")
	if key == "" {
		panic("env KEY required")
	}
	return &Config{
		Key: key,
	}
}

func (cfg *Config) GetCipherKey() string {
	return cfg.Key
}
