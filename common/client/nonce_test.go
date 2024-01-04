package client_test

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/ethereum-kits/common/client"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/fenghaojiang/ethereum-kits/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestOnListenNonce(t *testing.T) {
	t.Run("test on listen nonce", func(t *testing.T) {
		cli, err := client.NewEthereumClient("https://rpc.ankr.com/eth")
		require.Nil(t, err)
		require.NotNil(t, cli)

		addrNonce := model.NewAddressNonce(common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"))
		client.StartListeningNonce([]*client.EthereumClient{cli}, addrNonce, 1*time.Second)
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				log.Info("test listen nonce", zap.String("address", addrNonce.Address.String()), zap.Uint64("nonce", addrNonce.Nonce.Load()))
			}
		}

	})
}
