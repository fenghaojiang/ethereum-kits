package client

import (
	"context"
	"errors"
	"net/http"

	"github.com/fenghaojiang/ethereum-kits/common/log"
	"github.com/fenghaojiang/ethereum-kits/model"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type BundlerClient struct {
	cli *resty.Client
}

func NewBundlerClient() *BundlerClient {
	cli := resty.New()
	return &BundlerClient{
		cli: cli,
	}
}

func (b BundlerClient) SendBundle(ctx context.Context, endpoint string, headers map[string]string, request model.RPCRequest) (string, error) {
	var result model.RPCResponse[model.EthBundleResponse]
	resp, err := b.cli.R().SetHeaders(headers).
		SetBody(request).SetResult(&result).
		Post(endpoint)
	if err != nil {
		log.Error("failed to send bundle", zap.Error(err))
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		log.Error("failed to send bundle", zap.Int("statusCode", resp.StatusCode()), zap.Any("body", resp.Body()))
		return "", errors.New("failed to send bundle")
	}

	return result.Result.BundleHash, nil
}
