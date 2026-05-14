package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/liquentlabs/sdk-go/client/common"
	"github.com/liquentlabs/sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchange.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch LQT burnt details
	lqtBurntResponse, err := exchangeClient.FetchLqtBurnt(ctx)
	if err != nil {
		fmt.Printf("Failed to fetch LQT burnt details: %v\n", err)
		return
	}

	// Print JSON representation of the response
	jsonResponse, err := json.MarshalIndent(lqtBurntResponse, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal response to JSON: %v\n", err)
		return
	}

	fmt.Println("LQT Burnt Details:")
	fmt.Println(string(jsonResponse))
}
