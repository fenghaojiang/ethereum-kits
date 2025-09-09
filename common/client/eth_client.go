package client

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"go.uber.org/zap"
)

type EthereumClient struct {
	gethClient *gethclient.Client
	ethclient  *ethclient.Client
	Endpoint   string
}

func NewEthereumClient(endpoint string) (*EthereumClient, error) {
	rpcClient, err := rpc.DialContext(context.Background(), endpoint)
	if err != nil {
		log.Error("failed to create ethereum client", zap.Error(err))
		return nil, err
	}

	return &EthereumClient{
		gethClient: gethclient.New(rpcClient),
		ethclient:  ethclient.NewClient(rpcClient),
		Endpoint:   endpoint,
	}, nil
}

func (e *EthereumClient) ETHClient() *ethclient.Client {
	return e.ethclient
}

func (e *EthereumClient) GETHClient() *gethclient.Client {
	return e.gethClient
}

func (e *EthereumClient) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	err := e.ethclient.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Error("failed to send transaction", zap.Error(err))
		return err
	}

	return nil
}
