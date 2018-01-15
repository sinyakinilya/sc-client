package ethereum_client

import (
	"encoding/json"
	"fmt"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/model"
	"strconv"
)

func (ec EthereumClient) GetLastBlockNumber() (uint64, error) {
	var response model.GethRPCResponse
	byteResponse, err := ec.SendRequest("eth_blockNumber", nil)

	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return 0, err
	}

	lastBlock, err := strconv.ParseInt(response.Result.(string), 0, 64)

	if err != nil {
		return 0, err
	}

	return uint64(lastBlock), nil
}

func (ec EthereumClient) GetBlockByNumber(number uint64) (model.EthereumBlock, error) {
	var (
		params   []interface{}
		response model.GethRPCResponseBlock
	)
	params = append(params, fmt.Sprintf("0x%x", number), true)
	byteResponse, err := ec.SendRequest("eth_getBlockByNumber", params)

	if err != nil {
		return model.EthereumBlock{}, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		fmt.Println(string(byteResponse))
		return model.EthereumBlock{}, err
	}

	return response.Result, nil
}
