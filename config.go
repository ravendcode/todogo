package main

import (
	"os"

	"github.com/nicksnyder/go-i18n/i18n"
)

// Config App
type Config struct {
	Env     string
	Port    string
	DbPath  string
	Locales []string
}

// NewConfig func
func NewConfig() *Config {
	i18n.MustLoadTranslationFile("locale/en.yaml")
	i18n.MustLoadTranslationFile("locale/ru.yaml")

	config := &Config{
		Env:     os.Getenv("ENV"),
		DbPath:  "storage/data.db",
		Port:    "80",
		Locales: []string{"ru", "en"},
	}
	if os.Getenv("PORT") != "" {
		config.Port = os.Getenv("PORT")
	}
	if os.Getenv("ENV") != "development" {
		config.Env = "production"
	}
	return config
}
