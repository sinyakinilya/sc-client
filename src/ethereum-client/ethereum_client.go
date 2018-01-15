package ethereum_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/model"
	"github.com/suntechsoft/dmarket-go-backend/shared/util/env"
	"io/ioutil"
	"net/http"
)

type EthereumClient struct {
	HttpClient *http.Client
}

var (
	GethUrlJsonRpc = env.GetVar("ETHEREUM_NODE_ADDRESS", "http://127.0.0.1:8545")
	SCAddress      = env.GetVar("SC_ADDRESS", "0x0c328e06a90654bed74ab238983e3d883059b3af")
	SCOwnerAddress = env.GetVar("SC_OWNER_ADDRESS", "0x978bee7fbf556ca89fc542022c19b54a8662e501")
)

func (ec EthereumClient) SendRequest(methodName string, params []interface{}) ([]byte, error) {
	body := model.GethRPCRequest{
		JsonRPC: "2.0",
		Method:  methodName,
		Id:      "1",
	}
	body.Params = params
	requestBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	response, err := ec.Request(requestBody)
	var errorResponse model.GethRPCResponseError

	if err := json.Unmarshal(response, &errorResponse); err != nil {
		return nil, err
	}

	if errorResponse.Error.Code != 0 || errorResponse.Error.Message != "" {
		errorMessage, _ := json.Marshal(errorResponse.Error)
		return nil, errors.New(string(errorMessage))
	}

	return response, nil
}

func (ec EthereumClient) Request(requestBody []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", GethUrlJsonRpc, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := ec.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	return responseBody, nil
}
