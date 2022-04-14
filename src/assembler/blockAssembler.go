package assembler

import (
	"github.com/ethereum-block-indexer-service/models"
)

func BlockAssembler(records []*models.Block) map[string]interface{} {

	result := map[string]interface{}{}
	if len(records) == 0 {
		return result
	}
	blocks := []map[string]interface{}{}

	for _, record := range records {
		blocks = append(blocks, map[string]interface{}{
			"block_num":   record.BlockNumber,
			"block_hash":  record.Hash,
			"block_time":  record.Timestamp,
			"parent_hash": record.ParentHash,
		})
	}

	result["blocks"] = blocks
	return result
}
