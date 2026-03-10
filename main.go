package main

import (
	"fmt"
	"sync"
)

type Transaction struct {
	CorrelationID int
	Datetime      string
	ContaOrigem   int64
	ContaDestino  int64
	Valor         float64
}

func main() {

	transacoes := []Transaction{
		{1, "09/09/2023 14:15:00", 938485762, 2147483649, 150},
		{2, "09/09/2023 14:15:05", 2147483649, 210385733, 149},
		{3, "09/09/2023 14:15:29", 347586970, 238596054, 1100},
		{4, "09/09/2023 14:17:00", 675869708, 210385733, 5300},
		{5, "09/09/2023 14:18:00", 238596054, 674038564, 1489},
		{6, "09/09/2023 14:18:20", 573659065, 563856300, 49},
		{7, "09/09/2023 14:19:00", 938485762, 2147483649, 44},
		{8, "09/09/2023 14:19:01", 573659065, 675869708, 150},
	}

	executor := NewTransferExecutor()

	var wg sync.WaitGroup

	for _, t := range transacoes {

		wg.Add(1)

		go func(tr Transaction) {
			defer wg.Done()

			executor.Transferir(
				tr.CorrelationID,
				tr.ContaOrigem,
				tr.ContaDestino,
				tr.Valor,
			)

		}(t)

	}

	wg.Wait()

	fmt.Println("Processamento finalizado")
}