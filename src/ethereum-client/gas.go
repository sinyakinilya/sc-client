package ethereum_client

import (
	"encoding/json"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/model"
	"strconv"
)

func (ec EthereumClient) GetGasPrice() (dec uint64, hex string, err error) {
	var response model.GethRPCResponse
	byteResponse, err := ec.SendRequest("eth_gasPrice", nil)
	if err != nil {
		return 0, "", err
	}
	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return 0, "", err
	}
	decGasPrice, err := strconv.ParseInt(response.Result.(string), 0, 64)

	if err != nil {
		return 0, "", err
	}

	return uint64(decGasPrice), response.Result.(string), nil
}

func (ec EthereumClient) EstimateGas(txData model.GethSendTransactionParams) (dec uint64, hex string, err error) {
	var (
		response model.GethRPCResponse
		params   []interface{}
	)
	params = append(params, txData)
	byteResponse, err := ec.SendRequest("eth_estimateGas", params)
	if err != nil {
		return 0, "", err
	}
	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return 0, "", err
	}
	needGas, err := strconv.ParseInt(response.Result.(string), 0, 64)

	if err != nil {
		return 0, "", err
	}

	return uint64(needGas), response.Result.(string), nil
}
