package main

import (
	"fmt"
	"github.com/sinyakinilya/sc-client/src/ethereum-client"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/helper"
	"math/big"
	"net/http"
)

func main() {
	TokenHolderAddress := "0x7cb08d516a72dcd0c55f39ae647d5d19aebf2ad9"
	EthereumClient := ethereum_client.EthereumClient{HttpClient: &http.Client{}}
	SCOwnerAddress := ethereum_client.SCOwnerAddress
	SCAddress := ethereum_client.SCAddress
	erc20 := helper.ERC20{}
	amount := new(big.Int)

	amount.SetString("100000000000000", 0)
	executeApprove := erc20.Approve(SCOwnerAddress, amount)

	txApproveParams, err := EthereumClient.CreateTxParams(TokenHolderAddress, SCAddress, "0x0", executeApprove)
	if err != nil {
		panic(err)
	}
	gas := new(big.Int)
	gas.SetString(txApproveParams.Gas, 0)
	gasPrice := new(big.Int)
	gasPrice.SetString(txApproveParams.GasPrice, 0)
	etherAmount := new(big.Int)
	etherAmount.Mul(gas, gasPrice)
	hexAmount := fmt.Sprintf("0x%x", etherAmount)
	tt := helper.Ether(etherAmount).Text('f', 18)
	fmt.Println(gas, gasPrice, etherAmount, hexAmount, tt)

	//txSendEtherParam, err := EthereumClient.CreateTxParams(SCOwnerAddress, TokenHolderAddress, hexAmount, "")
	//if err != nil {
	//	panic(err)
	//}

	//	txHash := "0x087e92bd65e741195fba681648682136ce0c37faa4096a1e0f3a6400ef95bdcf"
	// 0x087e92bd65e741195fba681648682136ce0c37faa4096a1e0f3a6400ef95bdcf
	//
	txHash, err := EthereumClient.PersonalSendTransaction(txApproveParams, "test_test")
	if err != nil {
		panic(err)
	}
	fmt.Println(txHash)
	//var i int8
	//for i = 0; i < 100; i++ {
	//	status, err := EthereumClient.GetTransactionStatus(txHash)
	//	if err != nil {
	//		panic(err)
	//	}
	//	if status {
	//		break
	//	}
	//}

	/**
	1 вычисляем сколько нужно эфира для выполнения транзакции approve
	2 начисляем эфир на адрес на который отправили токены
	3 выполняем транзакцию approve для базового кошелька
	4 выполняем транзакцию transferFrom базовым кошельком
	*/

}
