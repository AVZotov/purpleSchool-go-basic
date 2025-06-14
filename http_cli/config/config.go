package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	XMasterKey string
}

func NewEnvConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}
	xMasterKey := os.Getenv("X_MASTER_KEY")
	if xMasterKey == "" {
		panic("env X_MASTER_KEY required")
	}
	return &Config{
		XMasterKey: xMasterKey,
	}
}

func (cfg *Config) GetMasterKey() string {
	return cfg.XMasterKey
}
