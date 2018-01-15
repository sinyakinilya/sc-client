package model

type GethRPCResponseTransactionReceipt struct {
	Id      string                     `json:"id"`
	JsonRPC string                     `json:"jsonrpc"`
	Result  EthereumTransactionReceipt `json:"result"`
}

type EthereumTransactionReceipt struct {
	TransactionHash   string        `json:"transactionHash"`
	TransactionIndex  string        `json:"transactionIndex"`
	BlockNumber       string        `json:"blockNumber"`
	BlockHash         string        `json:"blockHash"`
	CumulativeGasUsed string        `json:"cumulativeGasUsed"`
	GasUsed           string        `json:"GasUsed"`
	ContractAddress   string        `json:"contractAddress"`
	Logs              []interface{} `json:"logs"`
	LogsBloom         string        `json:"logsBloom"`
	Status            string        `json:"status"`
}
