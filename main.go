package main

import (
	"context"
	"github.com/ethereum-block-indexer-service/src/crawler"
	"github.com/ethereum-block-indexer-service/src/helpers/logger"
	"github.com/ethereum-block-indexer-service/src/resources/gormdao"
	"github.com/ethereum-block-indexer-service/src/routers"
)

func main() {

	defer logger.Close()
	defer gormdao.Close()

	ctx := context.Background()
	go crawler.Worker(ctx, int64(0))
	routers.Run()
}
