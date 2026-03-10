package main

import "fmt"

type TransferExecutor struct {
	db *AcessoDados
}

func NewTransferExecutor() *TransferExecutor {

	return &TransferExecutor{
		db: NewAcessoDados(),
	}

}

func (t *TransferExecutor) Transferir(
	correlationID int,
	contaOrigem int64,
	contaDestino int64,
	valor float64,
) {

	origem := t.db.GetSaldo(contaOrigem)

	if origem == nil {
		fmt.Println("Conta origem inexistente")
		return
	}

	if origem.Saldo < valor {

		fmt.Printf(
			"Transacao numero %d foi cancelada por falta de saldo\n",
			correlationID,
		)

		return
	}

	destino := t.db.GetSaldo(contaDestino)

	if destino == nil {
		fmt.Println("Conta destino inexistente")
		return
	}

	origem.Saldo -= valor
	destino.Saldo += valor

	fmt.Printf(
		"Transacao numero %d foi efetivada com sucesso! Novos saldos: Conta Origem: %.2f | Conta Destino: %.2f\n",
		correlationID,
		origem.Saldo,
		destino.Saldo,
	)
}