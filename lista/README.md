# Lista de Compras (CLI)

Este exercicio cria um programa simples de linha de comando para gerenciar uma lista de itens com quantidade.
O objetivo e praticar structs, metodos, slices, leitura de entrada e validacao em Go.

## O que foi feito

- Criada a struct `Item` com `Nome` e `Quantidade`.
- Criada a struct `Lista` com slice de itens e contador de itens.
- Implementados metodos para:
  - Adicionar item com quantidade.
  - Remover uma quantidade especifica do item.
  - Exibir a lista com as quantidades.
  - Limpar a lista.
  - Sair do programa.
- Tratamento de entrada do usuario (conversao para inteiro e validacao de quantidade).
- Comparacao de nomes ignorando maiusculas/minusculas, acentos e sinais.

## Como funciona

O programa exibe um menu em loop e executa a opcao escolhida.
Ao adicionar um item, o usuario informa o nome e a quantidade.
Ao remover, o usuario informa o nome e a quantidade a ser removida.

Regras de remocao:

- Se a quantidade solicitada for maior que a disponivel, a remocao e negada.
- Se for igual, o item e removido do slice.
- Se for menor, apenas a quantidade e reduzida.

## Normalizacao de texto

Para permitir remocao mesmo com grafias diferentes (ex.: "Feijao" vs "Feijao" com acento),
foi criada uma funcao de normalizacao que:

- Converte para minusculas.
- Remove espacos do inicio/fim.
- Substitui letras acentuadas por equivalentes sem acento.
- Remove sinais e mantem apenas letras e numeros.

Assim, nomes digitados de formas diferentes ainda correspondem ao item cadastrado.

## Conceitos praticados

- Structs e metodos com receiver.
- Slices e operacoes de append/remocao.
- Leitura de entrada com `bufio.Reader`.
- Conversao de string para inteiro com `strconv.Atoi`.
- Validacao de dados.
- Uso de `strings` e `unicode`.

## Como executar

```bash
go run lista.go
```

## Sugestoes de melhoria

- Unificar itens iguais na adicao (somar quantidade se ja existir).
- Exibir mensagem detalhada quando a remocao falhar.
- Salvar e carregar a lista em arquivo.
