package ethereum_client

import (
	"bytes"
	"encoding/json"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/model"
	"io/ioutil"
	"net/http"
)

type EthereumClient struct {
	JsonRpcUrl string
	HttpClient *http.Client
}

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
	if err != nil {
		return nil, err
	}
	var errorResponse model.GethRPCResponseError

	if err := json.Unmarshal(response, &errorResponse); err != nil {
		return nil, err
	}

	if errorResponse.Error.Code != 0 || errorResponse.Error.Message != "" {
		return nil, errorResponse.Error
	}

	return response, nil
}

func (ec EthereumClient) Request(requestBody []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", ec.JsonRpcUrl, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := ec.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	return responseBody, err
}
