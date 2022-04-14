package service

import (
	"context"
	"fmt"
	"github.com/ethereum-block-indexer-service/models"
	error2 "github.com/ethereum-block-indexer-service/src/error"
	"github.com/ethereum-block-indexer-service/src/helpers/config"
	"github.com/ethereum-block-indexer-service/src/helpers/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetNLatestBlock(ctx context.Context, n int64) (res []*models.Block, e *error2.Error) {

	log := logger.Logger()
	defer func() {
		if rcv := recover(); rcv != nil {
			log.Error("GetNLatestBlock panic", rcv)
			e = error2.UnexpectedError(rcv.(string))
		}
	}()

	res = make([]*models.Block, 0)
	jsonRpcEndPoint := config.Get("server.rpcUrl").(string)
	client, err := ethclient.DialContext(ctx, jsonRpcEndPoint)
	if err != nil {
		fmt.Println("err:", err)

	}
	defer client.Close()

	// Query the latest block
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	one := big.NewInt(1)
	inputNum := big.NewInt(n)
	end := block.Number().Sub(block.Number(), inputNum)

	for i := block.Number(); i.Cmp(end) > 0; i.Sub(i, one) {

		block, err := client.BlockByNumber(ctx, i)
		if err != nil {
			log.Fatal(err)
		}

		// Build the response to our model
		blocktemp := models.Block{
			BlockNumber: block.Number().Int64(),
			Timestamp:   block.Time(),
			Hash:        block.Hash().String(),
			ParentHash:  block.ParentHash().String(),
		}
		fmt.Println(len(block.Uncles()))
		res = append(res, &blocktemp)
		fmt.Printf("%+v\n", blocktemp)
	}

	return res, nil
}
