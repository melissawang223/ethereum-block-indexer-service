package service

import (
	"context"
	"github.com/ethereum-block-indexer-service/models"
	error2 "github.com/ethereum-block-indexer-service/src/error"
	"github.com/ethereum-block-indexer-service/src/helpers/config"
	"github.com/ethereum-block-indexer-service/src/helpers/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTxWithLogs(ctx context.Context, txHash string) (trans *models.Transaction, receipt *types.Receipt, e *error2.Error) {

	log := logger.Logger()
	defer func() {
		if rcv := recover(); rcv != nil {
			log.Error("GetTxWithLogs panic", rcv)
			e = error2.UnexpectedError("panic" + rcv.(string))
		}
	}()

	trans = &models.Transaction{}
	jsonRpcEndPoint := config.Get("server.rpcUrl").(string)
	client, err := ethclient.DialContext(ctx, jsonRpcEndPoint)
	if err != nil {
		log.Error("err:", err)
		return nil, nil, error2.EthClientError(err.Error())
	}
	defer client.Close()

	// Query the tx hash block
	hash := common.HexToHash(txHash)

	tx, pending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Error("err:", err)
		return nil, nil, error2.EthClientError(err.Error())
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()), nil)
	if err != nil {
		return nil, nil, error2.InternalServerError("signer error: " + err.Error())
	}

	// Transaction
	to := ""
	if msg.To() != nil {
		to = msg.To().String()
	}

	trans = &models.Transaction{
		Hash:     tx.Hash().String(),
		Value:    msg.From().String(),
		Gas:      tx.Gas(),
		GasPrice: msg.GasPrice().Uint64(),
		To:       to,
		Pending:  pending,
		Nonce:    tx.Nonce(),
		From:     msg.From().Hex(),
		Data:     msg.Data(),
	}

	receipt, err = client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.Fatal(err)
		return nil, nil, error2.EthClientError(err.Error())
	}

	return trans, receipt, nil
}
