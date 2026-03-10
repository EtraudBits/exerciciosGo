# EOF em Go (Guia Didatico)

Neste exercicio, o foco e entender o `EOF` durante leitura de dados em Go.

Observacao importante:

- EOF significa **End Of File** (fim do conteudo para leitura).
- No contexto deste exercicio, ele marca o fim do texto da pagina simulada.

## Objetivo do exercicio

Demonstrar como ler um conteudo em blocos pequenos (1 byte por vez) e decidir corretamente:

1. Quando finalizar a leitura com sucesso (`io.EOF`)
2. Quando tratar um erro real (`err != nil`)

Arquivo principal:

- `endoffire/endoffire.go`

## O que foi usado e por que

### `strings.NewReader(...)`

Foi usado para simular uma fonte de dados (como se fosse o corpo de uma pagina HTML), sem depender de rede.

```go
leitor := strings.NewReader("<html>Link1, Link2</html>")
```

Por que usar:

- Facilita estudar leitura de stream sem complexidade de HTTP.
- Mantem o foco no tratamento de EOF.

### `make([]byte, 1)`

Foi criado um buffer de 1 byte para ler caractere por caractere.

```go
p := make([]byte, 1)
```

Por que usar:

- Deixa visivel o ciclo de leitura.
- E didatico para mostrar quando o fim do conteudo acontece.

### `Read(p)` dentro de `for`

Leitura continua ate nao haver mais dados.

```go
_, err := leitor.Read(p)
```

Por que usar:

- `Read` e a API base de leitura em Go.
- Retorna erro para sinalizar fim (`io.EOF`) ou falha real.

### `if err == io.EOF` (estrategia central)

No codigo, EOF e tratado primeiro.

```go
if err == io.EOF {
	fmt.Println("\n -- Fim da Leitura: Pagina processada com sucesso --")
	break
}
```

Por que usar:

- `io.EOF` **nao e erro catastrofico**.
- E o sinal esperado de que todos os dados foram lidos.
- `break` encerra o loop de forma limpa.

### `if err != nil` para erro real

Depois de checar EOF, o codigo trata erros de verdade:

```go
if err != nil {
	fmt.Printf("Erro catastrofico na leitura: %v\n", err)
	return
}
```

Por que usar:

- Se nao for EOF e houver erro, algo realmente falhou.
- O `return` evita continuar em estado inconsistente.

### `fmt.Printf("%c", p[0])`

Quando `err` e `nil`, o byte lido e processado como caractere.

Por que usar:

- Mostra visualmente o processamento progressivo da entrada.

## Fluxo mental correto com EOF

1. Ler o proximo bloco com `Read`.
2. Se `err == io.EOF`, finalizar com sucesso.
3. Se `err != nil`, tratar como falha real.
4. Se `err == nil`, processar os dados lidos.

## Resultado esperado na execucao

Ao rodar:

```bash
go run endoffire.go
```

Saida esperada (conceitualmente):

1. Impressao do conteudo caractere por caractere
2. Mensagem final de leitura concluida ao atingir EOF

## Aprendizado principal

Neste exercicio, a estrategia importante e diferenciar:

- **fim normal da leitura** (`io.EOF`)
- **erro real de execucao** (`err != nil` diferente de EOF)

Essa separacao evita falsos alarmes e torna o codigo de leitura robusto e previsivel.
