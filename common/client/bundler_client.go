package client

import (
	"context"
	"encoding/json"
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
	resp, err := b.cli.R().SetHeaders(headers).
		SetBody(request).
		Post(endpoint)
	if err != nil {
		log.Error("failed to send bundle", zap.Any("resp", string(resp.Body())), zap.Error(err))
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		log.Error("failed to send bundle", zap.Int("statusCode", resp.StatusCode()), zap.Any("body", string(resp.Body())))
		return "", errors.New("failed to send bundle")
	}

	var result model.RPCResponse[model.EthBundleResponse]
	err = json.Unmarshal(resp.Body(), &result)
	if err == nil {
		return result.Result.BundleHash, nil
	}

	var strResult model.RPCResponse[string]
	err = json.Unmarshal(resp.Body(), &strResult)
	if err != nil {
		log.Error("failed to send bundle", zap.Any("body", string(resp.Body())), zap.Error(err))
		return "", err
	}

	return strResult.Result, nil
}
