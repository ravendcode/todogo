package todong

import (
	"os"
)

// Config App
type Config struct {
	Port string
}

var config *Config

func init() {
	config = &Config{Port: "80"}
	if os.Getenv("PORT") != "" {
		config.Port = os.Getenv("PORT")
	}
}

// GetConfig get Config instance
func GetConfig() *Config {
	return config
}
