package crawler

import (
	"context"
	"encoding/hex"
	"github.com/ethereum-block-indexer-service/models"
	"github.com/ethereum-block-indexer-service/src/helpers/config"
	"github.com/ethereum-block-indexer-service/src/helpers/logger"
	"github.com/ethereum-block-indexer-service/src/resources/gormdao"
	"github.com/ethereum-block-indexer-service/src/resources/gormdao/blockdao"
	"github.com/ethereum-block-indexer-service/src/resources/gormdao/eventLogsdao"
	"github.com/ethereum-block-indexer-service/src/resources/gormdao/transactiondao"
	"github.com/ethereum-block-indexer-service/src/service"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
)

var IndexRouter = make([]int64, 0)

var existMap ExistMap

type ExistMap struct {
	Data sync.Map
	Lock sync.RWMutex
}

// cached block that already in DB to avoid waste resource
func init() {

	existMap = ExistMap{
		Data: sync.Map{},
		Lock: sync.RWMutex{},
	}

	db := gormdao.DB()
	query := blockdao.NewQueryOption()
	blocks := blockdao.Gets(db, &query)

	existMap.Lock.Lock()
	for _, b := range *blocks {
		existMap.Data.Store(b.BlockNumber, true)
	}
	existMap.Lock.Unlock()
}

func Worker(ctx context.Context, blockNumber int64) {

	log := logger.Logger()
	defer func() {
		if r := recover(); r != nil {
			log.Error("Worker ERR %v", r)
		}
	}()

	var wg sync.WaitGroup
	head := blockNumber

	defer func() {
		if head != blockNumber {
			Worker(ctx, head)
		}
	}()

	jsonRpcEndPoint := config.Get("server.rpcUrl").(string)
	client, err := ethclient.DialContext(ctx, jsonRpcEndPoint)
	if err != nil {
		log.Error("ethclient Dial err:", err)
		return
	}
	defer client.Close()

	// Query the latest block
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumberH := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumberH)
	if err != nil {
		log.Fatal(err)
		return
	}

	head = block.Number().Int64()
	temp := config.Get("server.IndexerNumber").(int)
	indexLength := int64(temp)
	IndexRouter = make([]int64, indexLength)

	allLength := block.Number().Int64() - blockNumber
	length := allLength / indexLength

	for idx, _ := range IndexRouter {
		IndexRouter[idx] = blockNumber + length*int64(idx+1)
		wg.Add(1)
		go crawler(ctx, &wg, IndexRouter[idx], indexLength)
	}

	wg.Wait()
}

func crawler(ctx context.Context, wg *sync.WaitGroup, blockNumber, idxLength int64) {

	defer wg.Done()
	log := logger.Logger()
	jsonRpcEndPoint := config.Get("server.rpcUrl").(string)
	ethClient, err := ethclient.DialContext(ctx, jsonRpcEndPoint)
	if err != nil {
		log.Error("ethclient Dial err: %v", err)
		return
	}
	defer ethClient.Close()

	block, err := ethClient.BlockByNumber(ctx, big.NewInt(blockNumber))
	if err != nil {
		log.Fatal("BlockByNumber err:", err)
		return
	}

	db := gormdao.DB()
	one := big.NewInt(1)
	st := big.NewInt(blockNumber - idxLength)
	for i := block.Number(); i.Cmp(st) >= 0; i.Sub(i, one) {

		block, err := service.GetBlockById(ctx, i.Int64())
		//block, err := ethClient.BlockByNumber(ctx, i)
		if err != nil {
			log.Fatal("GetBlockById err:", err)
			continue
		}

		// Build the response to our model
		blocktemp := models.BlockModel{
			BlockNumber:       block.BlockNumber,
			ParentHash:        block.ParentHash,
			Difficulty:        block.Difficulty,
			Hash:              block.Hash,
			TransactionsCount: block.TransactionsCount,
			Timestamp:         block.Timestamp,
		}
		existMap.Lock.Lock()
		if _, ok := existMap.Data.Load(blocktemp.BlockNumber); !ok {

			errN := blockdao.New(db, &blocktemp)
			if errN != nil {
				log.Error("block new to db err %+v\n", errN.Error())
				continue
			}

			for _, tx := range block.Transactions {
				trans, receipt, err := service.GetTxWithLogs(ctx, tx.Hash)
				if err != nil {
					log.Error("GetTxWithLogs err %+v\n", err.Message())
					continue
				}
				tx.BlockNumber = block.BlockNumber
				// to has been decode int trans
				tx.To = trans.To

				//avoid nil data
				if tx.Data == nil {
					tx.Data = make([]byte, 0)
				}
				errN := transactiondao.New(db, &tx)
				if errN != nil {
					log.Error("transaction new to db err %+v\n", errN.Error())
					continue
				}
				if receipt != nil {
					for _, val := range receipt.Logs {
						encodedString := hex.EncodeToString(val.Data)
						eventLogsTemp := models.EventLogsModel{
							BlockNumber: val.BlockNumber,
							Hash:        val.TxHash.String(),
							Index:       val.Index,
							Data:        encodedString,
						}

						errEL := eventLogsdao.New(db, &eventLogsTemp)
						if errEL != nil {
							log.Error("event logs new to db err %+v\n", errN.Error())
							continue
						}
					}
				}
			}

			existMap.Data.Store(blocktemp.BlockNumber, true)

		}
		existMap.Lock.Unlock()
	}
	return
}
