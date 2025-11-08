package api

import (
	"algoflow-backend/internal/algorand"
	"algoflow-backend/internal/config"
	"encoding/json"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"net/http"
)

type Router struct {
	mux  *http.ServeMux
	port string
}

func NewRouter(cfg *config.Config, client *algod.Client) *Router {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong 🏓"))
	})

	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		status, err := algorand.GetNodeStatus(client)
		if err != nil {
			http.Error(w, "Failed to get Node status", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(status)
	})

	return &Router{
		mux:  mux,
		port: ":" + cfg.Port,
	}
}

func (r *Router) ListenAndServe() error {
	return http.ListenAndServe(r.port, r.mux)
}
