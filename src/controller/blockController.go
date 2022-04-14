package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/ethereum-block-indexer-service/src/assembler"
	"github.com/ethereum-block-indexer-service/src/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetLatestNBlockchain(ctx *gin.Context) {

	type rule struct {
		Limit int64 `valid:"required"`
	}

	limit, _ := ctx.GetQuery("limit")

	params := &rule{
		Limit: cast.ToInt64(limit),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "limit not valid, it is required",
		})
		return
	}

	// get blocks from rpc api
	blocks, err := service.GetNLatestBlock(ctx, params.Limit)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Message(),
		})
		return
	}

	// assemble the blocks to certain format
	result := assembler.BlockAssembler(blocks)

	ctx.JSON(200, result)
	return
}
