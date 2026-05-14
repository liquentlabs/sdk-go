package main

import (
	"context"
	"encoding/json"
	"fmt"

	chainstreamv2 "github.com/liquentlabs/sdk-go/chain/stream/types/v2"
	"github.com/liquentlabs/sdk-go/client"
	chainclient "github.com/liquentlabs/sdk-go/client/chain"
	"github.com/liquentlabs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		"",
		nil,
	)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint)

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	subaccountId := "0xbdaedec95d563fb05240d6e01821008454c24c36000000000000000000000000"

	lqtUsdtMarket := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	lqtUsdtPerpMarket := "0x17ef48032cb24375ba7c2e39f384e56433bcab20cbee9a7357e4cba2eb00abe6"

	req := chainstreamv2.StreamRequest{
		BankBalancesFilter: &chainstreamv2.BankBalancesFilter{
			Accounts: []string{"*"},
		},
		SpotOrdersFilter: &chainstreamv2.OrdersFilter{
			MarketIds:     []string{lqtUsdtMarket},
			SubaccountIds: []string{subaccountId},
		},
		DerivativeOrdersFilter: &chainstreamv2.OrdersFilter{
			MarketIds:     []string{lqtUsdtPerpMarket},
			SubaccountIds: []string{subaccountId},
		},
		SpotTradesFilter: &chainstreamv2.TradesFilter{
			MarketIds:     []string{lqtUsdtMarket},
			SubaccountIds: []string{"*"},
		},
		SubaccountDepositsFilter: &chainstreamv2.SubaccountDepositsFilter{
			SubaccountIds: []string{subaccountId},
		},
		DerivativeOrderbooksFilter: &chainstreamv2.OrderbookFilter{
			MarketIds: []string{lqtUsdtPerpMarket},
		},
		SpotOrderbooksFilter: &chainstreamv2.OrderbookFilter{
			MarketIds: []string{lqtUsdtMarket},
		},
		PositionsFilter: &chainstreamv2.PositionsFilter{
			SubaccountIds: []string{subaccountId},
			MarketIds:     []string{lqtUsdtPerpMarket},
		},
		DerivativeTradesFilter: &chainstreamv2.TradesFilter{
			SubaccountIds: []string{"*"},
			MarketIds:     []string{lqtUsdtPerpMarket},
		},
		OraclePriceFilter: &chainstreamv2.OraclePriceFilter{
			Symbol: []string{"LQT", "USDT"},
		},
	}

	ctx := context.Background()

	stream, err := chainClient.ChainStreamV2(ctx, req)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			res, err := stream.Recv()
			if err != nil {
				panic(err)
			}
			str, _ := json.MarshalIndent(res, "", "\t")
			fmt.Print(string(str))
		}
	}
}
