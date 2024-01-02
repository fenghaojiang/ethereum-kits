package constant_test

import (
	"math/big"
	"testing"

	"github.com/fenghaojiang/ethereum-kits/common/constant"
	"github.com/stretchr/testify/require"
)

func TestOnNumberOfETHFloat(t *testing.T) {
	t.Run("test on eth float", func(t *testing.T) {
		actual := constant.NumberOfETHFloat(0.1)

		require.Equal(t, new(big.Int).SetUint64(1_0000_0000_0000_0000_0), actual)

		actual = constant.NumberOfETHFloat(0.5)
		require.Equal(t, new(big.Int).SetUint64(5_0000_0000_0000_0000_0), actual)
	})
}
