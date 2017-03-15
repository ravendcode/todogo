package main

import (
	"os"
)

// Config App
type Config struct {
	ENV    string
	Port   string
	DbPath string
}

// NewConfig func
func NewConfig() *Config {
	config := &Config{ENV: os.Getenv("ENV"), Port: "80", DbPath: "storage/data.db"}
	if os.Getenv("PORT") != "" {
		config.Port = os.Getenv("PORT")
	}
	if os.Getenv("ENV") != "development" {
		config.ENV = "production"
	}
	return config
}
