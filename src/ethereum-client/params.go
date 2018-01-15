package ethereum_client

import (
	"github.com/sinyakinilya/sc-client/src/ethereum-client/model"
)

func (ec EthereumClient) CreateTxParams(from string, to string, value string, data string) (model.GethSendTransactionParams, error) {
	_, hexGasPrice, err := ec.GetGasPrice()
	if err != nil {
		return model.GethSendTransactionParams{}, err
	}

	params := model.GethSendTransactionParams{
		From:     from,
		To:       to,
		Value:    value,
		Gas:      "0x0", // hex from (1000000)
		GasPrice: hexGasPrice,
		Data:     data,
	}
	_, hexGas, err := ec.EstimateGas(params)
	if err != nil {
		return params, err
	}

	params.Gas = hexGas

	return params, nil
}
