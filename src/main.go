package main

import (
	"fmt"
	"github.com/sinyakinilya/sc-client/src/ethereum-client"
	"github.com/sinyakinilya/sc-client/src/ethereum-client/helper"
	"math/big"
	"net/http"
	"strings"
)

type TokenTransfer struct {
	from        string
	to          string
	amountToken *big.Int
	txHah       string
}

func main() {
	EthereumClient := ethereum_client.EthereumClient{HttpClient: &http.Client{}}
	erc20 := helper.ERC20{}
	lastBlock, _ := EthereumClient.GetLastBlockNumber()
	lastBlock = 4929009

	for lastBlock > 4929007 {
		fmt.Println("Current Block ", lastBlock)
		txs, _ := EthereumClient.GetTransactionFromBlock(lastBlock)
		transfers := []TokenTransfer{}
		for _, tx := range txs {
			if strings.ToLower(tx.To) != ethereum_client.SCAddress {
				continue
			}

			_, sc, err := EthereumClient.GetBuyerInfo(tx, ethereum_client.SCAddress)
			if err != nil {
				fmt.Println(err, sc)
			}

			// check tx status
			txStatus, err := EthereumClient.GetTransactionStatus(tx.Hash)
			if err != nil {
				panic(err)
			}
			if !txStatus && err == nil {
				continue
			}

			// check transfer signature
			to, amount, err := erc20.ParseTransferData(tx.Input)
			if err != nil && err.Error() != "input is not transfer data" {
				panic(err)
			}
			if err == nil {
				transfers = append(transfers, TokenTransfer{from: tx.From, to: to, amountToken: amount, txHah: tx.Hash})
				continue
			}

			// check transferFrom signature
			from, to, amount, err := erc20.ParseTransferFromData(tx.Input)
			if err != nil && err.Error() != "input is not transferFrom data" {
				panic(err)
			}
			if err == nil {
				transfers = append(transfers, TokenTransfer{from: from, to: to, amountToken: amount, txHah: tx.Hash})
			}



		}

		for _, transfer := range transfers {
			// TODO: проверить если транзакция со статусом "pending" с адресом "value.Params[0]" и кол токенов "value.Params[1]"
			// если есть то закрыть в "done" и выполнить перевод на базовый кошелек
			fmt.Println("TxHash: ", transfer.txHah, "From: ", transfer.from, " To:", transfer.to, " amount:", transfer.amountToken)
		}

		lastBlock--
	}
}
