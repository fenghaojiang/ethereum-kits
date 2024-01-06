package client_test

import (
	"testing"
	"time"

	"github.com/fenghaojiang/ethereum-kits/common/client"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestOnListeningBlock(t *testing.T) {
	t.Run("test on listen latest block", func(t *testing.T) {
		cli, err := client.NewEthereumClient("https://rpc.ankr.com/eth")
		require.Nil(t, err)
		require.NotNil(t, cli)

		client.StartListenLatestBlock([]*client.EthereumClient{
			cli,
		}, time.Second*1)

		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				log.Info("test latest block", zap.Uint64("latest", client.LatestBlock()))
			}
		}
	})
}
