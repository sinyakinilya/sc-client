package model

type ScMethod struct {
	FunctionName string
	Amount       uint64
	To           string
	Params       []interface{}
}