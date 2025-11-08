package config

import (
	"os"
)

type Config struct {
	Port       string
	AlgodAddr  string
	AlgodToken string
}

func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "8081"),                        // backend port
		AlgodAddr:  getEnv("ALGOD_ADDR", "http://127.0.0.1:8080"), // node port
		AlgodToken: getEnv("ALGOD_TOKEN", "634f0c3859a505d41ed73e68ce29191183276470c29e8b375e552ee308f50fee"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
