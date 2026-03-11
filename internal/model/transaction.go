package model

type Transaction struct {
	CorrelationID int
	Datetime      string
	ContaOrigem   int64
	ContaDestino  int64
	Valor         float64
}