package algorand

import (
	"context"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
)

// NodeStatus represents simplified status info
type NodeStatus struct {
	LastRound uint64 `json:"last_round"`
}

// GetNodeStatus fetches node status
func GetNodeStatus(client *algod.Client) (*NodeStatus, error) {
	status, err := client.Status().Do(context.Background())
	if err != nil {
		return nil, err
	}

	return &NodeStatus{
		LastRound: status.LastRound,
	}, nil
}
