package model

type EthBundleRequest struct {
	Txs               []string `json:"txs,omitempty"`
	BlockNumber       string   `json:"blockNumber,omitempty"`
	MinTimestamp      uint64   `json:"minTimestamp,omitempty"`
	MaxTimestamp      uint64   `json:"maxTimestamp,omitempty"`
	RevertingTxHashes []string `json:"revertingTxHashes,omitempty"`
	ReplacementUuid   string   `json:"replacementUuid,omitempty"`
}

type EthBundleResponse struct {
	BundleHash string `json:"bundleHash"`
}
