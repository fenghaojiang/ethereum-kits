package account_test

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fenghaojiang/ethereum-kits/common/account"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestOnGenerateNewAccount(t *testing.T) {
	priKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(priKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	log.Info("privateKeyHex", zap.Any("privateKeyHex", privateKeyHex))
	account, err := account.NewAccountFromPrivateKey(privateKeyHex)
	require.Nil(t, err)

	log.Info("account", zap.Any("account", account))

	publicKey := priKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	_ = crypto.FromECDSAPub(publicKeyECDSA)
	address1 := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	log.Info("address", zap.Any("address", address1))
}
