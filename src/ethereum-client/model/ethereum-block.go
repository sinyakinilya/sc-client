package model

type GethRPCResponseBlock struct {
	Id      string        `json:"id"`
	JsonRPC string        `json:"jsonrpc"`
	Result  EthereumBlock `json:"result"`
}

type EthereumBlockUncle string

type EthereumBlock struct {
	Difficulty       string                `json:"difficulty"`
	ExtraData        string                `json:"extraData"`
	GasLimit         string                `json:"gasLimit"`
	GasUsed          string                `json:"gasUsed"`
	Hash             string                `json:"hash"`
	LogsBloom        string                `json:"logsBloom"`
	Miner            string                `json:"miner"`
	MixHash          string                `json:"mixHash"`
	Nonce            string                `json:"nonce"`
	Number           string                `json:"number"`
	ParentHash       string                `json:"parentHash"`
	ReceiptsRoot     string                `json:"receiptsRoot"`
	Sha3Uncles       string                `json:"sha3Uncles"`
	Size             string                `json:"size"`
	StateRoot        string                `json:"stateRoot"`
	Timestamp        string                `json:"timestamp"`
	TotalDifficulty  string                `json:"totalDifficulty"`
	Transactions     []EthereumTransaction `json:"transactions"`
	TransactionsRoot string                `json:"transactionsRoot"`
	Uncles           []EthereumBlockUncle  `json:"uncles"`
}
