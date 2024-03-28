package account_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fenghaojiang/ethereum-kits/common/account"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestOnSignatureAuth(t *testing.T) {
	pk, addr, err := account.GenerateAccount()
	log.Info("new addr", zap.Any("account address", addr))
	require.Nil(t, err)

	acc, err := account.NewAccountFromPrivateKey(pk)
	require.Nil(t, err)

	t.Run("test on sign a message", func(t *testing.T) {
		sigHash := accounts.TextHash([]byte("test msg of ethereum kits"))
		sig, err := crypto.Sign(sigHash, acc.PrivateKey)
		require.Nil(t, err)

		resp_addr, err := crypto.SigToPub(sigHash, sig)

		require.Nil(t, err)
		require.Equal(t, addr, crypto.PubkeyToAddress(*resp_addr).Hex())
	})
}
