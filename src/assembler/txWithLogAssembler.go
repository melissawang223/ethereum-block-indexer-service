package assembler

import (
	"encoding/hex"
	"github.com/ethereum-block-indexer-service/models"
	"github.com/ethereum/go-ethereum/core/types"
)

func TxWithLogsAssembler(record *models.Transaction, receipt *types.Receipt) map[string]interface{} {

	encodedString := hex.EncodeToString(record.Data) //TODO
	result := map[string]interface{}{
		"tx_hash": record.Hash,
		"from":    record.From,
		"to":      record.To,
		"nonce":   record.Nonce,
		"data":    encodedString,
		"value":   record.Value,
	}

	transTx := []map[string]interface{}{}

	if receipt != nil {
		for _, val := range receipt.Logs {
			encodedString = hex.EncodeToString(val.Data) //TODO
			transTx = append(transTx, map[string]interface{}{
				"index": val.Index,
				"data":  encodedString,
			})
		}
	}

	result["logs"] = transTx
	return result
}
