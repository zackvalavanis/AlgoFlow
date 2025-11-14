package main

import (
	"algoflow-backend/internal/algorand"
	"algoflow-backend/internal/api"
	"algoflow-backend/internal/config"
	"log"
)

func main() {
	cfg := config.Load()

	// Connect to Algonode (MainNet) Indexer
	algodClient := algorand.NewAlgodClient("https://mainnet-api.algonode.cloud", "")
	indexerClient := algorand.NewIndexerClient("https://mainnet-idx.algonode.cloud", "")

	router := api.NewRouter(cfg, algodClient, indexerClient)
	log.Printf("🚀 Server running on port %s", cfg.Port)
	router.ListenAndServe()
}
