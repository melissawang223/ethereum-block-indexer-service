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

func GetBlockById(ctx context.Context, id int64) (res *models.BlockWithTx, e *error2.Error) {

	log := logger.Logger()
	defer func() {
		if rcv := recover(); rcv != nil {
			log.Error("GetBlockById panic", rcv)
			e = error2.UnexpectedError(rcv.(string))
		}
	}()

	res = &models.BlockWithTx{}
	jsonRpcEndPoint := config.Get("server.rpcUrl").(string)
	client, err := ethclient.DialContext(ctx, jsonRpcEndPoint)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer client.Close()

	// Query the latest block
	idBigInt := big.NewInt(id)
	header, _ := client.HeaderByNumber(context.Background(), idBigInt)
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	// Build the response to our model
	res = &models.BlockWithTx{
		BlockNumber:       block.Number().Uint64(),
		Timestamp:         block.Time(),
		Hash:              block.Hash().String(),
		ParentHash:        block.ParentHash().String(),
		Difficulty:        block.Difficulty().Uint64(),
		TransactionsCount: len(block.Transactions()),
		Transactions:      []models.Transaction{},
	}

	for _, tx := range block.Transactions() {
		t := models.Transaction{
			Hash:     tx.Hash().String(),
			Value:    tx.Value().String(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice().Uint64(),
			Nonce:    tx.Nonce(),
		}
		if tx.To() != nil {
			t.To = tx.To().String()
		}
		res.Transactions = append(res.Transactions, t)
	}

	return res, nil
}
