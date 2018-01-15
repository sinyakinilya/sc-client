package helper

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type ERC20 struct {
}

func (ERC20) Transfer(to string, amount *big.Int) string {
	return fmt.Sprintf("0xa9059cbb%064s%064x", to[2:], amount)
}

func (ERC20) TransferFrom(from string, to string, amount *big.Int) string {
	return fmt.Sprintf("0x23b872dd%064s%064s%064x", from[2:], to[2:], amount)
}

func (ERC20) Approve(to string, amount *big.Int) string {
	return fmt.Sprintf("0x095ea7b3%064s%064x", to[2:], amount)
}

func (ERC20) ParseTransferData(input string) (to string, amount *big.Int, err error) {
	//0xa9059cbb0000000000000000000000005237bc08b2fe644487366e246741bd7ec0eb24710000000000000000000000000000000000000000000000000000000005f5e100
	if strings.Index(input, "0xa9059cbb") != 0 {
		return to, amount, errors.New("input is not transfer data")
	}
	to = "0x" + input[34:74]
	amount = new(big.Int)
	amount.SetString(input[74:], 16)
	if !amount.IsUint64() {
		return to, amount, errors.New("bad amount data")
	}
	return to, amount, nil
}

func (ERC20) ParseTransferFromData(input string) (from string, to string, amount *big.Int, err error) {
	//0x23b872dd0000000000000000000000005237bc08b2fe644487366e246741bd7ec0eb24710000000000000000000000005237bc08b2fe644487366e246741bd7ec0eb24710000000000000000000000000000000000000000000000000000000005f5e100
	if strings.Index(input, "0x23b872dd") != 0 {
		return from, to, amount, errors.New("input is not transferFrom data")
	}
	from = "0x" + input[34:74]
	to = "0x" + input[98:138]
	amount = new(big.Int)
	amount.SetString(input[138:], 16)
	if !amount.IsUint64() {
		return from, to, amount, errors.New("bad amount data")
	}
	return from, to, amount, nil
}
