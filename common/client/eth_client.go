package client

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"go.uber.org/zap"
)

type EthereumClient struct {
	ethClient *ethclient.Client
	endpoint  string
}

func NewEthereumClient(endpoint string) (*EthereumClient, error) {
	ethClient, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Error("failed to create ethereum client", zap.Error(err))
		return nil, err
	}

	return &EthereumClient{
		ethClient: ethClient,
		endpoint:  endpoint,
	}, nil
}

func (e *EthereumClient) ETHClient() *ethclient.Client {
	return e.ethClient
}

func (e *EthereumClient) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	err := e.ethClient.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Error("failed to send transaction", zap.Error(err))
		return err
	}

	return nil
}
