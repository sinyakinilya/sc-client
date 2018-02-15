package helper

import (
	"math/big"
)

func Ether(wie *big.Int) (ether *big.Float) {
	var a, b big.Float
	a.SetInt(wie)
	b.SetInt(big.NewInt(1000000000000000000))

	return new(big.Float).Quo(&a, &b)
}

func Gwei(wei *big.Int) (ether *big.Float) {
	var a, b big.Float
	a.SetInt(wei)
	b.SetInt(big.NewInt(1000000000))

	return new(big.Float).Quo(&a, &b)
}
