package account

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/fenghaojiang/ethereum-kits/model"
	"go.uber.org/zap"
)

func (account *Account) GetAuthentication(request model.RPCRequest) (string, error) {
	b, err := json.Marshal(request)
	if err != nil {
		log.Error("failed to marshal request", zap.Error(err))
		return "", err
	}

	hashedBody := crypto.Keccak256Hash(b).Hex()
	sig, err := crypto.Sign(accounts.TextHash([]byte(hashedBody)), account.PrivateKey)
	if err != nil {
		log.Error("failed to sign to hashed body", zap.Any("request body", request), zap.Error(err))
		return "", err
	}

	return crypto.PubkeyToAddress(account.PrivateKey.PublicKey).Hex() + ":" + hexutil.Encode(sig), nil
}
