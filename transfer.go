package main

import (
	"fmt"
	"sync"
)

type TransferExecutor struct {
	db *AcessoDados
	mutex sync.Mutex
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

	t.mutex.Lock()
	defer t.mutex.Unlock()

	origem := t.db.GetSaldo(contaOrigem)

	if origem.Saldo < valor {

		fmt.Printf(
			"Transacao numero %d foi cancelada por falta de saldo\n",
			correlationID,
		)

		return
	}

	destino := t.db.GetSaldo(contaDestino)

	origem.Saldo -= valor
	destino.Saldo += valor

	fmt.Printf(
		"Transacao numero %d foi efetivada com sucesso! Novos saldos: Conta Origem: %.2f | Conta Destino: %.2f\n",
		correlationID,
		origem.Saldo,
		destino.Saldo,
	)
}