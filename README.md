# Transação Financeira - Refatoração e Correção de erros

## Descrição

Este projeto implementa uma simulação de transferências financeiras entre contas bancárias.
O código original apresentava alguns problemas de compilação, inconsistências de lógica e erros nas transações bancárias onde algumas estavam sendo canceladas mesmo com saldo positivo e outras sem saldo sendo efetivadas.

O objetivo deste projeto foi **refatorar a implementação original**, corrigir os problemas encontrados e aplicar boas práticas de engenharia de software.

A solução foi desenvolvida em **Golang**, mantendo o comportamento funcional da aplicação e melhorando sua arquitetura.

---

# Problemas identificados no código original

Durante a análise inicial do código foram encontrados os seguintes problemas:

### 1. Tipo incorreto para número de conta

O código utilizava o tipo `int` para representar números de conta.

Entretanto, alguns valores ultrapassavam o limite máximo permitido por `int`.

Exemplo:

2147483649

O limite de um `int` em sistemas de 32 bits é:

2147483647

Isso poderia causar erros de compilação.

✔ **Solução:**
O tipo foi alterado para `int64`, que suporta valores maiores.

---

### 2. Falta de validação de contas inexistentes

O código original não validava corretamente se as contas de origem ou destino existiam antes de acessar seus saldos.

Isso poderia gerar erros de execução como um NullPointer.

✔ **Solução:**
Foram adicionadas validações para garantir que as contas existam antes de processar a transferência.

---

### 3. Problema de concorrência

O sistema executa múltiplas transferências em paralelo.

No entanto, os saldos das contas eram atualizados simultaneamente por múltiplas threads sem qualquer mecanismo de sincronização.

Isso cria uma situação chamada **race condition**, onde duas ou mais operações leem e modificam o mesmo dado simultaneamente.

Por exemplo:

1. Duas transferências verificam o saldo da mesma conta ao mesmo tempo.
2. Ambas enxergam que há saldo suficiente.
3. As duas transferências são aprovadas.
4. O saldo é atualizado duas vezes.

Isso poderia gerar situações como:

* transações aprovadas sem saldo suficiente
* saldos negativos inesperados
* inconsistência de dados

Este problema pode ser identificado executando o programa com o detector de race do Go:

```
go run -race ./cmd
```

✔ **Solução:**

Foi adicionado um mecanismo de sincronização utilizando `sync.Mutex`, garantindo que apenas uma goroutine possa atualizar os saldos das contas por vez.  
Dessa forma, o saldo é verificado e atualizado de maneira segura, evitando problemas.

---

### 4. Muita responsabilidade em um arquivo só.

No código original, as classes eram responsáveis por múltiplas funções, como:

* acesso aos dados
* lógica de negócio
* manipulação de estado

Isso dificultava manutenção e testes.

✔ **Solução:**
O projeto foi refatorado aplicando **separação de responsabilidades**.

---

# Arquitetura da solução

Após a refatoração, o projeto foi organizado em camadas:

```
transacao-financeira

   main.go

internal/

   model/
      account.go
      transaction.go

   repository/
      account_repository.go

   service/
      transfer_service.go
```

### Camadas

**model**

Define as entidades do domínio da aplicação.

**repository**

Responsável pelo acesso e manipulação dos dados das contas.

**service**

Contém a lógica de negócio responsável por executar as transferências.

**main**

Ponto de entrada da aplicação.

---

# Passos realizados na refatoração

O processo de refatoração foi realizado em etapas:

1. Conversão do código original para Golang.
2. Correção de erros de compilação.
3. Ajuste dos tipos de dados utilizados.
4. Identificação de problemas de concorrência.
5. Implementação de controle de concorrência utilizando `mutex`.
6. Refatoração da arquitetura aplicando separação de responsabilidades.
7. Organização da estrutura do projeto seguindo boas práticas de Go.

---

# Como executar o projeto

Na raiz do projeto execute:

```
go mod tidy
```

Depois execute a aplicação:

```
go run . ou go run main.go
```

---

# Verificação de concorrência

Para verificar se há problemas de concorrência, execute:

```
go run -race .
```

Se nenhum alerta `DATA RACE` aparecer, significa que não há problemas de concorrência.

---

# Resultado

Com as melhorias aplicadas, o sistema agora possui:

* controle seguro de concorrência
* arquitetura mais organizada
* melhor separação de responsabilidades
* maior facilidade de manutenção

---


