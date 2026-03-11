package model

// representa uma conta bancária
type Account struct {
	ID      int64   // coloquei o id da conta int64 porque suporta numeros maiores evitando problemas.
	Saldo float64 //saldo da conta
}
