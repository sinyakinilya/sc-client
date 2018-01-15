package ethereum_client

import (
	"encoding/json"
	"math/big"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/model"
)

func (ec EthereumClient) GetBalance(address string, blockNumber string) (dec *big.Int, hex string, err error) {
	var (
		response model.GethRPCResponse
		params   []interface{}
	)
	dec = new(big.Int)
	params = append(params, address, blockNumber)
	byteResponse, err := ec.SendRequest("eth_getBalance", params)

	if err != nil {
		return dec, hex, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return dec, hex, err
	}

	dec.SetString(response.Result.(string), 0)

	return dec, response.Result.(string), nil
}

func (ec EthereumClient) CreateNewAccount(passphrase string) (address string, err error) {
	var (
		response model.GethRPCResponse
		params   []interface{}
	)
	params = append(params, passphrase)
	byteResponse, err := ec.SendRequest("personal_newAccount", params)

	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return "", err
	}

	return response.Result.(string), nil
}

func (ec EthereumClient) GetAccounts() (address []string, err error) {
	byteResponse, err := ec.SendRequest("eth_accounts", nil)

	if err != nil {
		return address, err
	}

	var response model.GethRPCResponseAccount

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return address, err
	}

	return response.Result, nil
}
