package model

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
)

type AddressNonce struct {
	Address common.Address
	nonce   *atomic.Uint64
}

func NewAddressNonce(address common.Address) *AddressNonce {
	return &AddressNonce{
		Address: address,
		nonce:   new(atomic.Uint64),
	}
}

func (a *AddressNonce) Nonce() uint64 {
	return a.nonce.Load()
}

func (a *AddressNonce) UpdatedNonce(newNonce uint64) {
	a.nonce.Store(newNonce)
}
