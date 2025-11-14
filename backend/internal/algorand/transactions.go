package algorand

import (
	"context"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/indexer"
	"log"
	"sort"
)

type TransactionSummary struct {
	AppID   uint64 `json:"app_id"`
	TxCount int    `json:"tx_count"`
}

// GetTopDapps fetches top dApps using the Indexer
func GetTopDapps(client *indexer.Client, txLimit int) ([]TransactionSummary, error) {
	appCounts := make(map[uint64]int)

	// Search for application call transactions
	resp, err := client.SearchForTransactions().
		TxType("appl").         // only ApplicationCall transactions
		Limit(uint64(txLimit)). // set the number of transactions to fetch
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	// Each tx in resp.Transactions is a flat struct
	for _, tx := range resp.Transactions {
		log.Printf("%+v\n", tx)
	}

	// Convert map to slice
	topDapps := make([]TransactionSummary, 0, len(appCounts))
	for appID, count := range appCounts {
		topDapps = append(topDapps, TransactionSummary{
			AppID:   appID,
			TxCount: count,
		})
	}

	// Sort descending by TxCount
	sort.Slice(topDapps, func(i, j int) bool {
		return topDapps[i].TxCount > topDapps[j].TxCount
	})

	// Only top 10
	if len(topDapps) > 10 {
		topDapps = topDapps[:10]
	}

	return topDapps, nil
}
