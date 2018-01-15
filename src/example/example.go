package main

import (
	"fmt"
	"github.com/sinyakinilya/sc-client/src/ethereum-client"
	"net/http"
)

func main() {
	EthereumClient := ethereum_client.EthereumClient{HttpClient: &http.Client{}}
	/*
		dec, gasPriceHex, err := EthereumClient.GetGasPrice()
		if err != nil {
			panic(err)
		}
		fmt.Println(dec, gasPriceHex)
	*/

	/*
		lastBlock, err := EthereumClient.GetLastBlockNumber()
		if err != nil {
			panic(err)
		}
		fmt.Println(lastBlock)
	*/

	/*
		block, err := EthereumClient.GetBlockByNumber(lastBlock)
		if err != nil {
			panic(err)
		}
		fmt.Println(block)
	*/

	/*
		txs, err := EthereumClient.GetTransactionFromBlock(lastBlock)
		if err != nil {
			panic(err)
		}
		fmt.Println(txs)
	*/

	/*
		hexNumber := fmt.Sprintf("0x%x", lastBlock)
		decBalance, hexBalance, err := EthereumClient.GetBalance("0x978bEE7FBF556CA89FC542022c19b54A8662E501", hexNumber)
		if err != nil {
			panic(err)
		}
		fmt.Println(decBalance, hexBalance)
	*/

	/*
		// 0x67c7ea46a692f7423bcb82f56b09fcb333f19f7b
		input := "0xa9059cbb0000000000000000000000005237bc08b2fe644487366e246741bd7ec0eb24710000000000000000000000000000000000000000000000000000000005f5e100"
		txParams, err := EthereumClient.CreateTxParams("0x978bEE7FBF556CA89FC542022c19b54A8662E501", ethereum_client.SCAddress, "0x0", input)
		if err != nil {
			panic(err)
		}
		fmt.Println(txParams)
	*/

	/*
		txHash, err := EthereumClient.PersonalSendTransaction(txParams, "...............")
		if err != nil {
			panic(err)
		}
		fmt.Println(txHash)
	*/

	/*
		address, err := EthereumClient.CreateNewAccount("test_test")
		if err != nil {
			panic(err)
		}
		fmt.Println(address)
	*/

	addresses, err := EthereumClient.GetAccounts()
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)

	txHash := "0x087e92bd65e741195fba681648682136ce0c37faa4096a1e0f3a6400ef95bdcf"
	txData, err := EthereumClient.GetTransactionByHash(txHash)
	if err != nil {
		panic(err)
	}
	fmt.Println(txData)

	status, err := EthereumClient.GetTransactionStatus(txHash)
	if err != nil {
		panic(err)
	}
	fmt.Println(status)

}
