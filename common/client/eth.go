package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fenghaojiang/ethereum-kits/common/constant"
	"github.com/samber/lo"
)

func (c *EthereumClient) BuildTransferEthLegacyTx(to common.Address, gasPrice, value *big.Int, nonce uint64) *types.Transaction {
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      constant.GasLimitOfTransferEth,
		To:       lo.ToPtr(to),
		Value:    value,
	})
}

// https://docs.alchemy.com/docs/maxpriorityfeepergas-vs-maxfeepergas
func (c *EthereumClient) BuildTransferEthDynamicTx(chainID *big.Int, to common.Address, maxPriorityFeePerGas, maxFeePerGas, value *big.Int, nonce uint64) *types.Transaction {
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: maxFeePerGas,
		Gas:       constant.GasLimitOfTransferEth,
		To:        lo.ToPtr(to),
		Value:     value,
	})
}
