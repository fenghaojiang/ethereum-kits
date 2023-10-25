package model

type RPCResponse[T any] struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Result  T      `json:"result"`
}
