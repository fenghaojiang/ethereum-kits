package client

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"go.uber.org/zap"
)

type Mempool struct {
	txnsLRU *expirable.LRU[string, *types.Transaction]
}

func NewMempool(size int, ttl time.Duration) *Mempool {
	return &Mempool{
		txnsLRU: expirable.NewLRU[string, *types.Transaction](size, func(key string, value *types.Transaction) {
			log.Info("transaction expired from lru", zap.String("txhash", key), zap.Any("transaction detail", value))
		}, ttl),
	}
}

func (m *Mempool) Add(tx *types.Transaction) {
	m.txnsLRU.Add(tx.Hash().String(), tx)
}

func (m *Mempool) MaxGasPriceTx() (legacyTx *types.Transaction, dynamicTx *types.Transaction) {
	for _, tx := range m.txnsLRU.Values() {
		if tx.Type() == types.LegacyTxType && legacyTx == nil {
			legacyTx = tx
			continue
		}

		if tx.Type() == types.DynamicFeeTxType && dynamicTx == nil {
			dynamicTx = tx
			continue
		}

		if tx.Type() == types.LegacyTxType && tx.GasPrice().Uint64() > legacyTx.GasPrice().Uint64() {
			legacyTx = tx
		}

		if tx.Type() == types.DynamicFeeTxType && tx.GasTipCap().Uint64() > dynamicTx.GasPrice().Uint64() {
			dynamicTx = tx
		}
	}

	return legacyTx, dynamicTx
}
