package client

import (
	"context"
	"time"

	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/fenghaojiang/ethereum-kits/model"
	"go.uber.org/zap"

	kitCommon "github.com/fenghaojiang/ethereum-kits/common"
)

func StartListeningNonce(ethereumClients []*EthereumClient, addressNonce *model.AddressNonce, watchInterval time.Duration) {
	if addressNonce == nil {
		log.Fatal("nonce should not be nil")
	}

	address := addressNonce.Address

	go kitCommon.Recoverable(context.Background(), func(ctx context.Context) {
		ticker := time.NewTicker(watchInterval)
		defer func() {
			ticker.Stop()
			log.Info("stop listening nonce of address", zap.String("address", address.String()))
		}()

		for {
			select {
			case <-ticker.C:
				log.Info("detecting latest nonce of address", zap.String("address", address.String()))
				for _, ec := range ethereumClients {
					nonce, err := ec.ETHClient().NonceAt(ctx, address, nil)
					if err != nil {
						log.Error("failed to get latest nonce of addr", zap.String("rpc endpoint", ec.endpoint), zap.Error(err))
						continue
					}

					log.Info("latest nonce of address fetched", zap.String("rpc endpoint", ec.endpoint), zap.String("address", address.String()), zap.Uint64("nonce", nonce))

					if nonce > addressNonce.Nonce.Load() {
						addressNonce.Nonce.Store(nonce)
						log.Info("latest nonce of address updated", zap.String("rpc endpoint", ec.endpoint), zap.String("address", address.String()), zap.Uint64("nonce", addressNonce.Nonce.Load()))
					}
				}

			}
		}
	})
}
