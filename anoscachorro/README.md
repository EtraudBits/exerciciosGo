## Anos de cachorro

Este programa le o nome e a idade humana de um cachorro e calcula a idade equivalente em anos de cachorro usando a regra 15/9/5:

- Primeiro ano vale 15 anos humanos.
- Segundo ano vale 9 anos humanos.
- A partir do terceiro ano, cada ano vale 5 anos humanos.

### O que foi feito

- Leitura do nome com validacao para aceitar apenas letras e espacos.
- Leitura da idade como numero decimal para permitir valores como 2.5.
- Calculo da idade canina usando a regra 15/9/5 com suporte a decimais.
- Mensagens de erro e repeticao do prompt quando a entrada e invalida.

### Pacotes usados

- `bufio`: leitura de entrada com buffer.
- `fmt`: entrada e saida formatada.
- `os`: acesso a entrada padrao.
- `strconv`: conversao de texto para numero decimal com `ParseFloat`.
- `strings`: limpeza da entrada com `TrimSpace`.
- `unicode`: validacao de letras e espacos no nome.

### Como executar

No terminal, dentro da pasta do exercicio:

```bash
go run anoscachorro.go
```

### Exemplo de uso

Entrada:

- Nome: `Bidu`
- Idade: `2.5`

Saida:

- `A idade do cao Bidu, e de 29.50 anos de cao`
