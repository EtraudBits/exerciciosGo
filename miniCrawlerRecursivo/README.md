# Mini Crawler Recursivo (Go)

Este exercício foi criado para praticar **recursividade** em Go simulando o comportamento de um crawler simples, que "visita" links em diferentes níveis de profundidade.

## Objetivo do exercício

Entender como uma função pode chamar a si mesma de forma controlada, com:

- **Caso base**: condição de parada para evitar loop infinito.
- **Passo recursivo**: nova chamada da função reduzindo o problema.
- **Ponto de retorno**: momento em que as chamadas começam a voltar na pilha.

## O que foi usado

- Linguagem **Go**
- Pacote padrão **fmt** para saída no terminal (`fmt.Printf`)
- Funções com parâmetros (`url string`, `profundidade int`)
- Conceito de **stack de chamadas** (cada chamada recursiva adiciona um novo frame)

## O que foi feito no programa

O programa define a função `explorarSite(url, profundidade)`, que simula uma navegação recursiva:

1. Verifica se a profundidade chegou a `0`.
   - Se chegou, imprime mensagem de limite atingido e finaliza essa chamada.
2. Se ainda há profundidade, imprime que está explorando a URL atual.
3. Gera um "próximo link" (concatenando `"/Link"`) e chama a própria função com `profundidade-1`.
4. Quando a chamada mais profunda termina, imprime a finalização da exploração da URL atual.

No `main`, a execução começa com:

```go
explorarSite("http://google.com", 3)
```

Isso produz uma sequência de exploração até o nível 0 e depois o retorno em ordem inversa.

## Aprendizados sobre recursividade

- Toda recursão precisa de **parada garantida**.
- Reduzir o problema a cada chamada (`profundidade-1`) é essencial.
- O retorno acontece de dentro para fora (última chamada é a primeira a retornar).
- Recursividade ajuda a modelar problemas hierárquicos (árvores, diretórios, links, etc.).

## Por que Pilha e não Lista?

Na recursão, o Go usa automaticamente a **pilha de chamadas** (stack), e isso combina com a lógica de exploração em profundidade:

- A pilha funciona em **LIFO** (_Last In, First Out_): a última chamada feita é a primeira a retornar.
- No crawler recursivo, isso significa: entra no link mais profundo primeiro e só depois volta para os anteriores.
- Cada chamada guarda seu próprio estado (`url` e `profundidade`) em um frame da pilha, sem precisar gerenciar manualmente.

Uma **lista** comum não resolve isso sozinha:

- Lista é uma estrutura genérica; para simular recursão com ela, você teria que implementar manualmente controle de inserção/remoção e ordem de processamento.
- Se usar lista como fila (FIFO), o comportamento muda para busca em largura (BFS), diferente da proposta deste exercício (profundidade/DFS).

Resumo: usamos pilha porque ela é o mecanismo natural da recursão e representa exatamente a ordem de ida e volta das chamadas.

### Exemplo visual da pilha de chamadas

Chamando:

```go
explorarSite("http://google.com", 3)
```

Montagem da pilha (ida):

```text
Topo
explorarSite(http://google.com/Link/Link/Link, 0)  <- caso base
explorarSite(http://google.com/Link/Link, 1)
explorarSite(http://google.com/Link, 2)
explorarSite(http://google.com, 3)
Base
```

Desmontagem da pilha (volta):

```text
1) retorna profundidade 0
2) retorna profundidade 1
3) retorna profundidade 2
4) retorna profundidade 3
```

Isso mostra exatamente o comportamento **LIFO** da pilha: o último nível que entrou é o primeiro a sair.

## Como executar

Dentro da pasta `miniCrawlerRecursivo`:

```bash
go run miniCrawlerRecursivo.go
```

## Sugestões de evolução

- Trocar links simulados por leitura real de páginas.
- Adicionar controle de URLs já visitadas para evitar repetição.
- Limitar também por quantidade máxima de páginas.
- Tratar erros de requisição e timeout.
