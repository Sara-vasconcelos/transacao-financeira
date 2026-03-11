package main

import (
	"sync"

	"transacao-financeira/internal/model"
	"transacao-financeira/internal/repository"
	"transacao-financeira/internal/service"
)

func main() {

	// cria repository de contas
	repo := repository.NewAccountRepository()

	// cria service
	executor := service.NewTransferService(repo)

	// lista de transações
	transacoes := []model.Transaction{
		{CorrelationID: 1, Datetime: "09/09/2023 14:15:00", ContaOrigem: 938485762, ContaDestino: 2147483649, Valor: 150},
		{CorrelationID: 2, Datetime: "09/09/2023 14:15:05", ContaOrigem: 2147483649, ContaDestino: 210385733, Valor: 149},
		{CorrelationID: 3, Datetime: "09/09/2023 14:15:29", ContaOrigem: 347586970, ContaDestino: 238596054, Valor: 1100},
		{CorrelationID: 4, Datetime: "09/09/2023 14:17:00", ContaOrigem: 675869708, ContaDestino: 210385733, Valor: 5300},
		{CorrelationID: 5, Datetime: "09/09/2023 14:18:00", ContaOrigem: 238596054, ContaDestino: 674038564, Valor: 1489},
		{CorrelationID: 6, Datetime: "09/09/2023 14:18:20", ContaOrigem: 573659065, ContaDestino: 563856300, Valor: 49},
		{CorrelationID: 7, Datetime: "09/09/2023 14:19:00", ContaOrigem: 938485762, ContaDestino: 2147483649, Valor: 44},
		{CorrelationID: 8, Datetime: "09/09/2023 14:19:01", ContaOrigem: 573659065, ContaDestino: 675869708, Valor: 150},
	}

	var wg sync.WaitGroup //serve para controlar múltiplas goroutines. Permite que o programa espere todas as goroutines terminarem antes de finalizar.

	for _, t := range transacoes {

		wg.Add(1) //uma nova goroutine será executada

		go func(tr model.Transaction) { //cria uma goroutine
			defer wg.Done() //finaliza e quando terminar ele avisa o waitGroup : essa tarefa terminou
			//chamando o serviço que executa a transferencia
			executor.Transferir(
				tr.CorrelationID,
				tr.ContaOrigem,
				tr.ContaDestino,
				tr.Valor,
			)

		}(t)

	}

	wg.Wait() //Aqui o programa fica bloqueado até todas as goroutines terminarem.

}
