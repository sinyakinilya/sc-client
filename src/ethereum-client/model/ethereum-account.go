package model

type GethRPCResponseAccount struct {
	Id      string   `json:"id"`
	JsonRPC string   `json:"jsonrpc"`
	Result  []string `json:"result"`
}
