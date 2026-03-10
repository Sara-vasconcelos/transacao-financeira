package main

type AcessoDados struct {
	TabelaSaldos []ContaSaldo
}

func NewAcessoDados() *AcessoDados {

	return &AcessoDados{
		TabelaSaldos: []ContaSaldo{
			{938485762, 180},
			{347586970, 1200},
			{2147483649, 0},
			{675869708, 4900},
			{238596054, 478},
			{573659065, 787},
			{210385733, 10},
			{674038564, 400},
			{563856300, 1200},
		},
	}
}

func (a *AcessoDados) GetSaldo(conta int64) *ContaSaldo {

	for i := range a.TabelaSaldos {

		if a.TabelaSaldos[i].Conta == conta {
			return &a.TabelaSaldos[i]
		}

	}

	return nil
}