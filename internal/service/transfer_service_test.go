package service

import (
	"testing"
	"transacao-financeira/internal/repository"
)

func TestTransferir_ComSucesso(t *testing.T) {

	repo := repository.NewAccountRepository()
	service := NewTransferService(repo)

	contaOrigem := int64(938485762)
	contaDestino := int64(347586970)

	origemAntes := repo.GetAccount(contaOrigem).Saldo
	destinoAntes := repo.GetAccount(contaDestino).Saldo

	valor := 50.0

	service.Transferir(1, contaOrigem, contaDestino, valor)

	origemDepois := repo.GetAccount(contaOrigem).Saldo
	destinoDepois := repo.GetAccount(contaDestino).Saldo

	if origemDepois != origemAntes-valor {
		t.Errorf("esperava saldo da conta origem %.2f, mas recebeu %.2f",
			origemAntes-valor,
			origemDepois,
		)
	}

	if destinoDepois != destinoAntes+valor {
		t.Errorf("esperava saldo da conta destino %.2f, mas recebeu %.2f",
			destinoAntes+valor,
			destinoDepois,
		)
	}
}

func TestTransferir_ContaOrigemInexistente(t *testing.T) {

	repo := repository.NewAccountRepository()
	service := NewTransferService(repo)

	contaOrigem := int64(999999999)
	contaDestino := int64(347586970)

	destinoAntes := repo.GetAccount(contaDestino).Saldo

	service.Transferir(1, contaOrigem, contaDestino, 50)

	destinoDepois := repo.GetAccount(contaDestino).Saldo

	if destinoAntes != destinoDepois {
		t.Errorf("o saldo da conta destino não deveria ser alterado")
	}
}

func TestTransferir_ContaDestinoInexistente(t *testing.T) {

	repo := repository.NewAccountRepository()
	service := NewTransferService(repo)

	contaOrigem := int64(938485762)
	contaDestino := int64(999999999)

	origemAntes := repo.GetAccount(contaOrigem).Saldo

	service.Transferir(1, contaOrigem, contaDestino, 50)

	origemDepois := repo.GetAccount(contaOrigem).Saldo

	if origemAntes != origemDepois {
		t.Errorf("o saldo da conta origem não deveria ser alterado")
	}
}

func TestTransferir_SaldoInsuficiente(t *testing.T) {

	repo := repository.NewAccountRepository()
	service := NewTransferService(repo)

	contaOrigem := int64(210385733) // conta com saldo baixo
	contaDestino := int64(347586970)

	origemAntes := repo.GetAccount(contaOrigem).Saldo
	destinoAntes := repo.GetAccount(contaDestino).Saldo

	service.Transferir(1, contaOrigem, contaDestino, 500)

	origemDepois := repo.GetAccount(contaOrigem).Saldo
	destinoDepois := repo.GetAccount(contaDestino).Saldo

	if origemAntes != origemDepois {
		t.Errorf("o saldo da conta origem não deveria mudar quando não há saldo suficiente")
	}

	if destinoAntes != destinoDepois {
		t.Errorf("o saldo da conta destino não deveria mudar quando a transferência falha")
	}
}