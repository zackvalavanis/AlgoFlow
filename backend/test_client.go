package main

import (
	"context"
	"fmt"
	"log"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
)

func main() {
	client, err := algod.MakeClient("http://127.0.0.1:8080", "634f0c3859a505d41ed73e68ce29191183276470c29e8b375e552ee308f50fee")
	if err != nil {
		log.Fatalf("Failed to make client: %s", err)
	}

	status, err := client.Status().Do(context.Background())
	if err != nil {
		log.Fatalf("Failed to get status: %s", err)
	}

	fmt.Printf("Node status: %+v\n", status)
}
