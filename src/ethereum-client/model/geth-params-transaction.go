package model

type GethSendTransactionParams struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data"`
}

type GethSendSmartContractParams struct {
	To       string `json:"to"`
	Data     string `json:"data"`
}