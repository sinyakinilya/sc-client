package main

import (
	"fmt"
	"github.com/sinyakinilya/sc-client/src/ethereum-client"
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
	erc20 := ethereum_client.ERC20{}
	lastBlock, _ := EthereumClient.GetLastBlockNumber()
	lastBlock = 2301917

	for lastBlock > 2301914 {
		fmt.Println("Current Block ", lastBlock)
		txs, _ := EthereumClient.GetTransactionFromBlock(lastBlock)
		transfers := []TokenTransfer{}
		for _, tx := range txs {
			if strings.ToLower(tx.To) != ethereum_client.SCAddress {
				continue
			}

			// check tx status
			txStatus, err := EthereumClient.GetTransactionStatus(tx.Hash)
			if !txStatus && err == nil {
				continue
			} else {
				panic(err)
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
