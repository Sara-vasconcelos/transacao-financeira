package repository

import (
	"testing"
)

func TestNewAccountRepository(t *testing.T) {

	repo := NewAccountRepository()

	if repo == nil {
		t.Errorf("esperava que o repositório fosse inicializado")
	}

	if len(repo.accounts) == 0 {
		t.Errorf("esperava que o repositório contivesse contas cadastradas")
	}
}

func TestGetAccount_WhenAccountExists(t *testing.T) {

	repo := NewAccountRepository()

	account := repo.GetAccount(938485762)

	if account == nil {
		t.Fatalf("esperava encontrar a conta, mas retornou nil")
	}

	if account.ID != 938485762 {
		t.Errorf("esperava ID da conta 938485762, mas recebeu %d", account.ID)
	}
}

func TestGetAccount_WhenAccountDoesNotExist(t *testing.T) {

	repo := NewAccountRepository()

	account := repo.GetAccount(999999999)

	if account != nil {
		t.Errorf("esperava nil para conta inexistente")
	}
}

func TestUpdateAccount(t *testing.T) {

	repo := NewAccountRepository()

	account := repo.GetAccount(938485762)

	account.Saldo= 999

	repo.UpdateAccount(account)

	updatedAccount := repo.GetAccount(938485762)

	if updatedAccount.Saldo != 999 {
		t.Errorf("esperava saldo 999, mas recebeu %f", updatedAccount.Saldo)
	}
}