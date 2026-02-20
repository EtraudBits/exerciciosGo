# Registro de Notas e Buscas (Go)

## Objetivo

Este programa de terminal permite cadastrar alunos, registrar notas, buscar informações, calcular média individual e persistir os dados em arquivo JSON.

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
- Persistência dos dados em `notas.json`.
- Carregamento automático dos dados ao iniciar o programa.
- Confirmação antes de ações sensíveis:
  - alterar notas
  - remover aluno
  - sair do sistema

### Persistência de dados

- `carregarDados()` é chamada no início do `main`.
- A função tenta ler `notas.json` com `os.ReadFile`.
- Se o arquivo existir e o conteúdo for válido, os alunos são carregados com `json.Unmarshal`.
- Após carregar, o mapa de índices (`map[string]int`) é reconstruído para manter buscas em O(1).
- `salvarDados()` é chamada sempre que a lista muda:
  - ao adicionar aluno
  - ao remover aluno
  - ao atualizar as 4 notas
  - ao incluir nota específica (1 a 4)

### Exemplo do arquivo `notas.json`

```json
[
  {
    "nome": "Maria",
    "notas": [8.5, 7.0, 9.0, 8.0]
  },
  {
    "nome": "Joao",
    "notas": [6.5, 7.5, 8.0, 7.0]
  }
]
```

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
- `encoding/json` para serialização e desserialização dos dados
- `fmt` para exibição de mensagens
- `os` para leitura/escrita de arquivo e encerramento do programa
- `strconv` para converter texto em número
- `strings` para tratar texto de entrada
- `unicode` para validar caracteres do nome

### Estruturas e conceitos

- `struct` para modelar `aluno` e `registroNotas`
- métodos com receiver para regras de negócio
- arrays fixos (`[4]float64`) para armazenar as 4 notas
- **`map[string]int`** para otimizar busca por nome (O(1) em vez de O(n))
- `for` e `switch` para fluxo do menu
- funções auxiliares para validação e confirmação
- persistência em JSON (`MarshalIndent`, `Unmarshal`, `ReadFile`, `WriteFile`)

### Otimização com Map

A estrutura `registroNotas` agora usa um mapa (`indices map[string]int`) que mapeia o nome do aluno para seu índice no slice. Isso permite:

- Busca por nome em tempo constante O(1)
- Atualização e remoção também otimizadas
- O mapa é sincronizado com o slice em cada operação (adicionar, remover)

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

Ao informar as notas de um aluno, o programa calcula corretamente a média individual com base nas 4 notas cadastradas/atualizadas e mantém os dados salvos em `notas.json` entre execuções.
