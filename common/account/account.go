package account

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"go.uber.org/zap"
)

type Account struct {
	PrivateKeyHex  string
	PrivateKey     *ecdsa.PrivateKey
	AccountAddress common.Address
}

func NewAccountFromPrivateKey(privateKeyHex string) (*Account, error) {
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Error("failed to load private key", zap.Error(err))
		return nil, err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &Account{
		PrivateKeyHex:  privateKeyHex,
		PrivateKey:     privateKey,
		AccountAddress: address,
	}, nil
}

func NewAccountFromMemonic(phrase string) (*Account, error) {
	seed, err := hdwallet.NewSeedFromMnemonic(phrase)
	if err != nil {
		log.Error("failed to generate seed from mnemonic", zap.Error(err))
		return nil, err
	}

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Error("failed to generate wallet from seed", zap.Error(err))
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Error("failed to derive account from wallet", zap.Error(err))
		return nil, err
	}

	privateKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Error("failed to get private key hex from wallet", zap.Error(err))
		return nil, err
	}

	return NewAccountFromPrivateKey(privateKeyHex)
}

func (a *Account) SignTx(unsignedTx *types.Transaction, s types.Signer) (*types.Transaction, error) {
	signedTx, err := types.SignTx(unsignedTx, s, a.PrivateKey)
	if err != nil {
		log.Error("failed to sign transaction", zap.Error(err))
		return nil, err
	}

	return signedTx, nil
}
