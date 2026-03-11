package repository

import "transacao-financeira/internal/model"

// representa um repositorio de contas
type AccountRepository struct {
	accounts map[int64]*model.Account // ex: 938485762 → {ID: 938485762, Balance: 180}
}

func NewAccountRepository() *AccountRepository {

	accounts := map[int64]*model.Account{ //map de contas já inicializado.
		938485762:  {ID: 938485762, Saldo: 180},
		347586970:  {ID: 347586970, Saldo: 1200},
		2147483649: {ID: 2147483649, Saldo: 0},
		675869708:  {ID: 675869708, Saldo: 4900},
		238596054:  {ID: 238596054, Saldo: 478},
		573659065:  {ID: 573659065, Saldo: 787},
		210385733:  {ID: 210385733, Saldo: 10},
		674038564:  {ID: 674038564, Saldo: 400},
		563856300:  {ID: 563856300, Saldo: 1200},
	}

	return &AccountRepository{
		accounts: accounts,
	}
}

// busca uma conta pelo ID.
func (r *AccountRepository) GetAccount(id int64) *model.Account {

	account, exists := r.accounts[id] //verifica se existe

	if !exists {
		return nil
	}

	return account
}

// atualiza uma conta no repository
func (r *AccountRepository) UpdateAccount(account *model.Account) {

	r.accounts[account.ID] = account

}
