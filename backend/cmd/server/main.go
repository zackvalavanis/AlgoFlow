package main

import (
	"algoflow-backend/internal/algorand"
	"algoflow-backend/internal/api"
	"algoflow-backend/internal/config"
	"log"
)

func main() {
	cfg := config.Load()

	client := algorand.NewClient("http://127.0.0.1:8080", "634f0c3859a505d41ed73e68ce29191183276470c29e8b375e552ee308f50fee")

	router := api.NewRouter(cfg, client)
	log.Printf("🚀 Server running on port %s", cfg.Port)
	router.ListenAndServe()
}
