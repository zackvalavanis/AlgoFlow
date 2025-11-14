package api

import (
	"algoflow-backend/internal/algorand"
	"algoflow-backend/internal/config"
	"encoding/json"
	"net/http"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/indexer"
)

type Router struct {
	mux           *http.ServeMux
	port          string
	algodClient   *algod.Client
	indexerClient *indexer.Client
}

func NewRouter(cfg *config.Config, algodClient *algod.Client, indexerClient *indexer.Client) *Router {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong 🏓"))
	})

	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		status, err := algorand.GetNodeStatus(algodClient)
		if err != nil {
			http.Error(w, "Failed to get Node status", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(status)
	})

	mux.HandleFunc("/api/top-dapps", func(w http.ResponseWriter, r *http.Request) {
		topDapps, err := algorand.GetTopDapps(indexerClient, 1000)
		if err != nil {
			http.Error(w, "Failed to fetch top dapps", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(topDapps)
	})

	return &Router{
		mux:           mux,
		port:          ":" + cfg.Port,
		algodClient:   algodClient,
		indexerClient: indexerClient,
	}
}

func (r *Router) ListenAndServe() error {
	return http.ListenAndServe(r.port, r.mux)
}
