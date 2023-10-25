package common

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"go.uber.org/zap"
)

func ToRawTransaction(tx *types.Transaction) (string, error) {
	data, err := tx.MarshalBinary()
	if err != nil {
		log.Error("failed to marshal transaction", zap.Error(err))
		return "", err
	}

	return hexutil.Encode(data), nil
}
