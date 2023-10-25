package account

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/fenghaojiang/ethereum-kits/common/log"
	"go.uber.org/zap"
)

func GenerateAccount() (string, string, error) {
	priKey, err := crypto.GenerateKey()
	if err != nil {
		log.Error("failed to generate private key", zap.Error(err))
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(priKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	// log.Info("privateKeyHex", zap.Any("privateKeyHex", privateKeyHex))
	_, err = NewAccountFromPrivateKey(privateKeyHex)
	if err != nil {
		log.Error("failed to create account", zap.Error(err))
		return "", "", err
	}

	// log.Info("account", zap.Any("account", account))

	publicKey := priKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("error casting public key to ECDSA")
	}

	_ = crypto.FromECDSAPub(publicKeyECDSA)
	address1 := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// log.Info("address", zap.Any("address", address1))

	return privateKeyHex, address1, nil
}
