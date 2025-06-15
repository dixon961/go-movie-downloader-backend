package config

import "os"

type Config struct {
	Port           string
	JackettBaseURL string
	JackettAPIKey  string
}

func Load() *Config {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:           port,
		JackettBaseURL: os.Getenv("JACKETT_BASE_URL"),
		JackettAPIKey:  os.Getenv("JACKETT_API_KEY"),
	}
}
