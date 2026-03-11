package model

// representa o modelo de uma transação financeira
type Transaction struct {
	CorrelationID int    //identificador da transação.
	Datetime      string //momento em que a transação aconteceu.
	ContaOrigem   int64
	ContaDestino  int64
	Valor         float64
}
