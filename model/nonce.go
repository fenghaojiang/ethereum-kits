package model

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
)

type AddressNonce struct {
	Address common.Address
	Nonce   *atomic.Uint64
}

func NewAddressNonce(address common.Address) *AddressNonce {
	return &AddressNonce{
		Address: address,
		Nonce:   new(atomic.Uint64),
	}
}
