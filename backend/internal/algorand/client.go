package algorand

import (
	"context"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"log"
)

// NewClient creates a new Algorand SDK client
func NewClient(addr, token string) *algod.Client {
	client, err := algod.MakeClient(addr, token)
	if err != nil {
		log.Fatalf("Failed to create Algorand client: %s", err)
	}
	return client
}

// GetNodeStatus fetches the current node status
func GetNodeStatus(client *algod.Client) (interface{}, error) {
	status, err := client.Status().Do(context.Background())
	if err != nil {
		return nil, err
	}
	return status, nil
}
