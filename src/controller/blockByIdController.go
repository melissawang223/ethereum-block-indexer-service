package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/ethereum-block-indexer-service/src/assembler"
	"github.com/ethereum-block-indexer-service/src/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetBlockById(ctx *gin.Context) {

	type rule struct {
		Id int64 `valid:"required"`
	}

	params := &rule{
		Id: cast.ToInt64(ctx.Param("id")),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "limit not valid, it is required",
		})
		return
	}

	// get blocks from rpc api
	blocks, err := service.GetBlockById(ctx, params.Id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Message(),
		})
		return
	}

	// assemble the blocks to certain format
	result := assembler.BlockWithTxAssembler(blocks)

	ctx.JSON(200, result)
	return
}
