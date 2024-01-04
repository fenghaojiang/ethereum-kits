package client

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/fenghaojiang/ethereum-kits/common"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"go.uber.org/zap"
)

var LatestBlock = new(atomic.Uint64)

func StartListenLatestBlock(ethereumClients []*EthereumClient, watchInterval time.Duration) {
	go common.Recoverable(context.Background(), func(ctx context.Context) {
		ticker := time.NewTicker(watchInterval)
		defer func() {
			ticker.Stop()
			log.Info("stop listening latest block")
		}()

		for {
			select {
			case <-ticker.C:
				log.Info("detecting latest block...")
				for _, ec := range ethereumClients {
					number, err := ec.ETHClient().BlockNumber(ctx)
					if err != nil {
						log.Error("failed to get latest block", zap.String("rpc endpoint", ec.endpoint), zap.Error(err))
						continue
					}

					log.Info("latest block fetched", zap.String("rpc endpoint", ec.endpoint), zap.Uint64("latest block", number))

					if number > LatestBlock.Load() {
						LatestBlock.Store(number)
						log.Info("latest block updated", zap.String("rpc endpoint", ec.endpoint), zap.Uint64("latest block", LatestBlock.Load()))
					}
				}
			}
		}
	})
}
