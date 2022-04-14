package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/ethereum-block-indexer-service/src/assembler"
	"github.com/ethereum-block-indexer-service/src/service"
	"github.com/gin-gonic/gin"
)

func GetTxWithLogs(ctx *gin.Context) {

	type rule struct {
		TxHash string `valid:"required"`
	}

	params := &rule{
		TxHash: ctx.Param("txHash"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "TxHash not valid, it is required",
		})
		return
	}

	// get blocks from rpc api
	trans, receipt, err := service.GetTxWithLogs(ctx, params.TxHash)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Message(),
		})
		return
	}

	// assemble the blocks to certain format
	result := assembler.TxWithLogsAssembler(trans, receipt)

	ctx.JSON(200, result)

	return
}
