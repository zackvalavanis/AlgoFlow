package algorand

import (
	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/indexer"
	"log"
)

// NewAlgodClient creates a new Algod client
func NewAlgodClient(address, token string) *algod.Client {
	client, err := algod.MakeClient(address, token)
	if err != nil {
		log.Fatalf("failed to create algod client: %v", err)
	}
	return client
}

// NewIndexerClient creates a new Indexer client
func NewIndexerClient(address, token string) *indexer.Client {
	client, err := indexer.MakeClient(address, token)
	if err != nil {
		log.Fatalf("failed to create indexer client: %v", err)
	}
	return client
}
