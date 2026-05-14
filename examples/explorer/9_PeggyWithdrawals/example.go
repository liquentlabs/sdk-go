package main

import (
	"context"
	"encoding/json"
	"fmt"

	explorerPB "github.com/liquentlabs/sdk-go/exchange/explorer_rpc/pb"

	"github.com/liquentlabs/sdk-go/client/common"
	explorerclient "github.com/liquentlabs/sdk-go/client/explorer"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	explorerClient, err := explorerclient.NewExplorerClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	sender := "lqt14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"

	req := explorerPB.GetPeggyWithdrawalTxsRequest{
		Sender: sender,
	}

	res, err := explorerClient.GetPeggyWithdrawals(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Print(string(str))
}
