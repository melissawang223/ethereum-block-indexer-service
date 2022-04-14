package assembler

import (
	"github.com/ethereum-block-indexer-service/models"
)

func BlockWithTxAssembler(record *models.BlockWithTx) map[string]interface{} {

	result := map[string]interface{}{
		"block_num":   record.BlockNumber,
		"block_hash":  record.Hash,
		"block_time":  record.Timestamp,
		"parent_hash": record.ParentHash,
	}

	transTx := []string{}

	for _, transaction := range record.Transactions {
		transTx = append(transTx, transaction.Hash)
	}

	result["transactions"] = transTx
	return result
}
