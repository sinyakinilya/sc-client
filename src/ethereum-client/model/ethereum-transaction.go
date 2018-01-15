package model

type GethRPCResponseTransaction struct {
	Id      string              `json:"id"`
	JsonRPC string              `json:"jsonrpc"`
	Result  EthereumTransaction `json:"result"`
}

type EthereumTransaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	To               string `json:"to"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"` // call smart-contract function
	Nonce            string `json:"nonce"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}
