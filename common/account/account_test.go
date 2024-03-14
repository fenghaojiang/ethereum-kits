package account_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/ethereum-kits/common/account"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const (
	privateKeyInHexForTest = "0x5fcd6ddc5471acbe3c0f48b43c5a0c4ca364b137a0340f01e38efbb85e7ce0ac"
	memonicPhraseForTest   = "ten able puppy cart gadget change health must creek human bracket above"

	accountAddressForTestFromPhrase = "0x688DB4fC60816A5838485d0ab9b4ef358f849c35"
	accountAddressForTestFromHex    = "0x81EbdaCabf245849412183f13cb83E9a7ffF2da4"
)

func TestOnNewAccount(t *testing.T) {
	t.Run("test on new account from private key", func(t *testing.T) {
		_account, err := account.NewAccountFromPrivateKey(privateKeyInHexForTest)
		if err != nil {
			log.Error("failed to create account", zap.Error(err))
			t.Fail()
		}

		log.Info("account", zap.Any("account", _account))

		require.True(t, account.ValidateAccount(_account, common.HexToAddress(accountAddressForTestFromHex)))

		_account1, err := account.NewAccountFromPrivateKey("0xd8dd45d24cd1d8ff776853e40fdf2e036c8bc516e62e5503437dfab3ab5685d3")
		if err != nil {
			log.Error("failed to create account", zap.Error(err))
			t.Fail()
		}

		log.Info("account1", zap.Any("account1", _account1))
	})

	t.Run("test on new account from memonic phrase", func(t *testing.T) {
		_account, err := account.NewAccountFromMemonic(memonicPhraseForTest, 0)
		if err != nil {
			log.Error("failed to create account", zap.Error(err))
			t.Fail()
		}

		log.Info("account", zap.Any("account", _account))
		_account1, err := account.NewAccountFromMemonic(memonicPhraseForTest, 1)
		if err != nil {
			log.Error("failed to create account", zap.Error(err))
			t.Fail()
		}

		log.Info("account1", zap.Any("account1", _account1))

		require.True(t, account.ValidateAccount(_account, common.HexToAddress(accountAddressForTestFromPhrase)))
	})
}
