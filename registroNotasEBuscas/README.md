# Registro de Notas e Buscas (Go)

## Objetivo

Este programa de terminal permite cadastrar alunos, registrar notas, buscar informações e calcular média individual.

## O que foi implementado

- Menu interativo com 8 opções:
  - 1.  Adicionar aluno
  - 2.  Buscar aluno
  - 3.  Atualizar as 4 notas
  - 4.  Calcular média do aluno
  - 5.  Remover aluno
  - 6.  Exibir alunos
  - 7.  Incluir nota (1 a 4) em aluno existente
  - 8.  Sair

- Cadastro de aluno com **4 notas fixas**.
- Cálculo de média individual por aluno usando:
  - soma das 4 notas
  - divisão por 4
- Busca de aluno por nome.
- Atualização completa das 4 notas de um aluno.
- Atualização de nota específica (nota 1, 2, 3 ou 4).
- Remoção de aluno.
- Confirmação antes de ações sensíveis:
  - alterar notas
  - remover aluno
  - sair do sistema

## Validações aplicadas

- Nome do aluno:
  - obrigatório
  - aceita apenas letras e espaços

- Nota:
  - precisa ser número válido
  - em caso inválido, exibe mensagem de erro

- Opção de nota na atualização pontual:
  - aceita somente valores entre 1 e 4

## O que foi usado no programa

### Linguagem

- Go

### Pacotes da biblioteca padrão

- `bufio` para leitura de entrada no terminal
- `fmt` para exibição de mensagens
- `os` para encerrar o programa
- `strconv` para converter texto em número
- `strings` para tratar texto de entrada
- `unicode` para validar caracteres do nome

### Estruturas e conceitos

- `struct` para modelar `aluno` e `registroNotas`
- métodos com receiver para regras de negócio
- arrays fixos (`[4]float64`) para armazenar as 4 notas
- `for` e `switch` para fluxo do menu
- funções auxiliares para validação e confirmação

## Como executar

No diretório `registroNotasEBuscas`:

```bash
go run registroNotasEBuscas.go
```

Opcional (compilar antes):

```bash
go build registroNotasEBuscas.go
```

## Resultado esperado

Ao informar as notas de um aluno, o programa calcula corretamente a média individual com base nas 4 notas cadastradas/atualizadas.
