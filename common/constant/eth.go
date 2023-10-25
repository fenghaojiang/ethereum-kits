package constant

import "math/big"

var (
	gwei = new(big.Int).SetUint64(1_00000_0000)
	eth  = new(big.Int).Mul(gwei, gwei)
)

func NumberOfGWei(number uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(number), gwei)
}

func GWei() *big.Int {
	return new(big.Int).Set(gwei)
}

func NumberOfETH(number uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(number), eth)
}

func ETH() *big.Int {
	return new(big.Int).Set(eth)
}

const (
	EthereumDecimals = 18
	USDCDecimals     = 6
)

// Gas
const (
	GasLimitOfTransferEth = 21_000
)
