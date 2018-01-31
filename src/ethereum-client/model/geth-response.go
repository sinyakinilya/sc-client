package model

import "encoding/json"

type GethRPCError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type GethRPCResponse struct {
	Id      string      `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

type GethRPCResponseError struct {
	Id      string       `json:"id"`
	JsonRPC string       `json:"jsonrpc"`
	Error   GethRPCError `json:"error"`
}

func (ge GethRPCError) Error() string {
	errorMessage, _ := json.Marshal(ge)
	return string(errorMessage)
}