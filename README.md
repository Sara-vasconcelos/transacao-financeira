## Identificação de problema de concorrência

O sistema executa múltiplas transferências em paralelo utilizando goroutines.

Durante a execução foi identificado um problema de **race condition** ao acessar e modificar os saldos das contas.

Isso ocorre porque múltiplas goroutines podem atualizar o saldo da mesma conta simultaneamente, gerando resultados inconsistentes.

Exemplo:

- duas transferências podem ler o mesmo saldo ao mesmo tempo
- ambas validam que existe saldo suficiente
- ambas debitem o valor

Resultado: saldo incorreto ou transações inválidas sendo aprovadas.

Esse comportamento pode ser observado executando o programa com o detector de race do Go:

```bash
go run -race .
```