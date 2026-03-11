package repository

import "transacao-financeira/internal/model"

type AccountRepository struct {
	accounts map[int64]*model.Account
}

func NewAccountRepository() *AccountRepository {

	accounts := map[int64]*model.Account{
		938485762: {ID: 938485762, Balance: 180},
		347586970: {ID: 347586970, Balance: 1200},
		2147483649: {ID: 2147483649, Balance: 0},
		675869708: {ID: 675869708, Balance: 4900},
		238596054: {ID: 238596054, Balance: 478},
		573659065: {ID: 573659065, Balance: 787},
		210385733: {ID: 210385733, Balance: 10},
		674038564: {ID: 674038564, Balance: 400},
		563856300: {ID: 563856300, Balance: 1200},
	}

	return &AccountRepository{
		accounts: accounts,
	}
}

func (r *AccountRepository) GetAccount(id int64) *model.Account {

	account, exists := r.accounts[id]

	if !exists {
		return nil
	}

	return account
}

func (r *AccountRepository) UpdateAccount(account *model.Account) {

	r.accounts[account.ID] = account

}