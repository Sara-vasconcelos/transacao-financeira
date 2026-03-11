package service

import (
	"fmt"
	"sync"

	"transacao-financeira/internal/repository"
)

// representa o serviço que executa transferências.
type TransferService struct {
	repo  *repository.AccountRepository //dependência da service, ele usa o repository para buscar contas e atualizar contas
	mutex sync.Mutex                    // É como uma trava que evita que varias goroutines alterarem o saldo ao mesmo tempo. Apenas uma transferência por vez possa modificar os saldos
}

// Nova instancia da service
func NewTransferService(repo *repository.AccountRepository) *TransferService {

	return &TransferService{
		repo: repo,
	}

}

// executa uma transferência entre duas contas.
func (t *TransferService) Transferir(
	correlationID int,
	contaOrigem int64,
	contaDestino int64,
	valor float64,
) {

	t.mutex.Lock()         //lock() bloqueia acesso concorrente.Somente uma goroutine pode executar essa parte do código
	defer t.mutex.Unlock() //Garante que o mutex será liberado quando a função terminar

	origem := t.repo.GetAccount(contaOrigem) //Busca a conta no repository

	//Se a conta não existir, a transferência é cancelada
	if origem == nil {
		fmt.Println("Conta origem inexistente")
		return
	}

	destino := t.repo.GetAccount(contaDestino) //Busca a conta destino.

	if destino == nil {
		fmt.Println("Conta destino inexistente") //se não encontrar , cancela
		return
	}

	if origem.Saldo < valor { //verifica se a  conta origem tem saldo suficiente.

		fmt.Printf(
			"Transacao numero %d foi cancelada por falta de saldo\n",
			correlationID,
		)

		return
	}

	origem.Saldo -= valor  //perde dinheiro
	destino.Saldo += valor //recebe dinheiro

	//atualiza os saldos  no repository
	t.repo.UpdateAccount(origem)
	t.repo.UpdateAccount(destino)

	fmt.Printf(
		"Transacao numero %d foi efetivada com sucesso! Novos saldos: Conta Origem: %.2f | Conta Destino: %.2f\n",
		correlationID,
		origem.Saldo,
		destino.Saldo,
	)
}
