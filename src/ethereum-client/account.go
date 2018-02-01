package ethereum_client

import (
	"encoding/json"
	"math/big"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/model"
	"fmt"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/helper"
	"strings"
	"errors"
)

func (ec EthereumClient) GetBalance(address string, blockNumber uint64) (dec *big.Int, hex string, err error) {
	var (
		response model.GethRPCResponse
		params   []interface{}
	)
	dec = new(big.Int)
	params = append(params, address, fmt.Sprintf("0x%x", blockNumber))
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

func (ec EthereumClient) GetBuyerInfo(tx model.EthereumTransaction, scAddress string) (buyer string, sc model.ScMethod, err error) {
	if strings.ToLower(tx.To) != scAddress {
		return buyer, sc, errors.New("this transaction is not related to the smart contract")
	}
	erc20 := helper.ERC20{}

	// check transfer signature
	to, amount, err := erc20.ParseTransferData(tx.Input)
	if err != nil && err.Error() != "input is not transfer data" {
		return buyer, sc, err
	}
	if err == nil {
		sc.FunctionName = "transfer"
		sc.Params = append(sc.Params, to, amount)
		sc.Amount = amount.Uint64()
		sc.To = to

		return tx.From, sc, nil
	}

	// check transferFrom signature
	from, to, amount, err := erc20.ParseTransferFromData(tx.Input)
	if err != nil {
		if err.Error() == "input is not transferFrom data" {
			err = errors.New("input data is not transfer/transferFrom call function")
		}
		return buyer, sc, err
	}

	sc.FunctionName = "transferFrom"
	sc.Params = append(sc.Params, from, to, amount)
	sc.Amount = amount.Uint64()
	sc.To = to

	return from, sc, nil
}

func (ec EthereumClient) GetTokenBalance(scAddress string, address string, blockNumber uint64) (dec *big.Int, hex string, err error) {
	var (
		response model.GethRPCResponse
		params   []interface{}
	)
	dec = new(big.Int)
	erc20 := helper.ERC20{}
	balanceOfData, err := erc20.GetBalanceOf(address)
	if err != nil {
		return dec, hex, err
	}
	scParams := model.GethSendSmartContractParams{
		To:scAddress,
		Data: balanceOfData,
	}

	params = append(params, scParams, fmt.Sprintf("0x%x", blockNumber))
	byteResponse, err := ec.SendRequest("eth_call", params)

	if err != nil {
		return dec, hex, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return dec, hex, err
	}

	dec.SetString(response.Result.(string), 0)

	return dec, response.Result.(string), nil
}

func (ec EthereumClient) GetAllowance(scAddress string, owner string, spender string, blockNumber uint64) (dec *big.Int, hex string, err error) {
	var (
		response model.GethRPCResponse
		params   []interface{}
	)
	dec = new(big.Int)
	erc20 := helper.ERC20{}
	allowanceData, err := erc20.GetAllowance(owner, spender)
	if err != nil {
		return dec, hex, err
	}
	scParams := model.GethSendSmartContractParams{
		To:scAddress,
		Data: allowanceData,
	}

	params = append(params, scParams, fmt.Sprintf("0x%x", blockNumber))
	byteResponse, err := ec.SendRequest("eth_call", params)

	if err != nil {
		return dec, hex, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return dec, hex, err
	}

	dec.SetString(response.Result.(string), 0)

	return dec, response.Result.(string), nil
}

func (ec EthereumClient) UnlockAccount(address string, passphrase string, second int) (ok bool, err error) {
	var (
		response model.GethRPCResponse
		params   []interface{}
	)
	params = append(params, address, passphrase, second)
	byteResponse, err := ec.SendRequest("personal_unlockAccount", params)
	if err != nil {
		switch e := err.(type) {
		case model.GethRPCError:
			if (e.Code == -32000) && (strings.Compare(e.Message, "could not decrypt key with given passphrase") == 0) {
				return false, nil
			} else {
				return ok, err
			}
		default:
			return ok, err
		}
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return false, err
	}

	return response.Result.(bool), nil
}
