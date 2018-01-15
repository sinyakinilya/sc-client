package ethereum_client

import (
	"encoding/json"
	"fmt"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/model"
)

func (ec EthereumClient) GetTransactionFromBlock(blockNumber uint64) ([]model.EthereumTransaction, error) {
	block, err := ec.GetBlockByNumber(blockNumber)

	if err != nil {
		return nil, err
	}

	return block.Transactions, nil
}

func (ec EthereumClient) PersonalSendTransaction(txData model.GethSendTransactionParams, passphrase string) (txHash string, err error) {
	var (
		params   []interface{}
		response model.GethRPCResponse
	)
	params = append(params, txData, passphrase)
	byteResponse, err := ec.SendRequest("personal_sendTransaction", params)

	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return "", err
	}

	return response.Result.(string), nil
}

func (ec EthereumClient) GetTransactionByHash(txHash string) (txData model.EthereumTransaction, err error) {
	var (
		response model.GethRPCResponseTransaction
		params   []interface{}
	)
	params = append(params, txHash)
	byteResponse, err := ec.SendRequest("eth_getTransactionByHash", params)
	if err != nil {
		return txData, err
	}
	fmt.Println(string(byteResponse))
	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return txData, err
	}

	return response.Result, nil
}

func (ec EthereumClient) GetTransactionReceipt(txHash string) (txData model.EthereumTransactionReceipt, err error) {
	var (
		response model.GethRPCResponseTransactionReceipt
		params   []interface{}
	)
	params = append(params, txHash)
	byteResponse, err := ec.SendRequest("eth_getTransactionReceipt", params)
	if err != nil {
		return txData, err
	}
	fmt.Println(string(byteResponse))
	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return txData, err
	}

	return response.Result, nil
}

func (ec EthereumClient) GetTransactionStatus(txHash string) (status bool, err error) {
	receipt, err := ec.GetTransactionReceipt(txHash)
	if err != nil {
		return false, err
	}

	if receipt.Status == "0x0" {
		status = false
	} else {
		status = true
	}

	return status, nil
}
