package model

// {
// 	"jsonrpc": "2.0",
// 	"id": 1,
// 	"method": "eth_sendBundle",
// 	"params": [
// 	  {
// 		txs,               // Array[String], A list of signed transactions to execute in an atomic bundle, list can be empty for bundle cancellations
// 		blockNumber,       // (Optional) String, a hex-encoded block number for which this bundle is valid. Default, current block number
// 		minTimestamp,      // (Optional) Number, the minimum timestamp for which this bundle is valid, in seconds since the unix epoch
// 		maxTimestamp,      // (Optional) Number, the maximum timestamp for which this bundle is valid, in seconds since the unix epoch
// 		revertingTxHashes, // (Optional) Array[String], A list of tx hashes that are allowed to revert or be discarded
// 		replacementUuid,   // (Optional) String, any arbitrary string that can be used to replace or cancel this bundle
// 		refundPercent,     // (Optional) Number, the percentage (from 0 to 99) of the  ETH reward of the last transaction, or the transaction specified by refundIndex, that should be refunded back to the ‘refundRecipient’
// 		refundIndex,       // (Optional) Number, the index of the transaction of which the ETH reward should be refunded. Default, last transaction in the bundle
// 		refundRecipient,   // (Optional) Address, the address that will receive the ETH refund. Default, sender of the first transaction in the bundle
// 	  }
// 	]
// }

type RPCRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

func CommonRPCRequest(method string, params []any) RPCRequest {
	return RPCRequest{
		JsonRPC: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}
}
