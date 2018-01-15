package helper

import "math/big"

func Ether(wie *big.Int) (ether *big.Float) {
	ether = new(big.Float)
	ether.SetString(wie.String())
	fwei := new(big.Float)
	fwei.SetString("0.000000000000000001")

	return ether.Mul(ether, fwei)
}
