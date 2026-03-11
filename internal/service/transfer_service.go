package service

import (
	"fmt"
	"sync"

	"transacao-financeira/internal/repository"
)

type TransferService struct {
	repo  *repository.AccountRepository
	mutex sync.Mutex
}

func NewTransferService(repo *repository.AccountRepository) *TransferService {

	return &TransferService{
		repo: repo,
	}

}

func (t *TransferService) Transferir(
	correlationID int,
	contaOrigem int64,
	contaDestino int64,
	valor float64,
) {

	t.mutex.Lock()
	defer t.mutex.Unlock()

	origem := t.repo.GetAccount(contaOrigem)

	if origem == nil {
		fmt.Println("Conta origem inexistente")
		return
	}

	destino := t.repo.GetAccount(contaDestino)

	if destino == nil {
		fmt.Println("Conta destino inexistente")
		return
	}

	if origem.Balance < valor {

		fmt.Printf(
			"Transacao numero %d foi cancelada por falta de saldo\n",
			correlationID,
		)

		return
	}

	origem.Balance -= valor
	destino.Balance += valor

	t.repo.UpdateAccount(origem)
	t.repo.UpdateAccount(destino)

	fmt.Printf(
		"Transacao numero %d foi efetivada com sucesso! Novos saldos: Conta Origem: %.2f | Conta Destino: %.2f\n",
		correlationID,
		origem.Balance,
		destino.Balance,
	)
}