## Churrasco

Este programa calcula a quantidade de carne necessaria para um churrasco e o custo total, com base no numero de pessoas e no preco do kg.

### O que foi feito

- Leitura do numero de pessoas e do preco do kg em reais.
- Conversao do preco de reais (texto) para centavos com uma funcao auxiliar.
- Calculo da quantidade total de carne usando a media de 500 g por pessoa.
- Conversao de gramas para kg para exibir o resultado.
- Calculo do custo total em centavos e exibicao formatada em reais.

### Pacotes usados

- `fmt`: entrada e saida formatada.
- `strconv`: conversao de string para inteiro.
- `strings`: limpeza e separacao do valor em reais.

### Como executar

No terminal, dentro da pasta do exercicio:

```bash
go run churrasco.go
```

### Exemplo de uso

Entrada:

- Pessoas: `6`
- Preco do kg: `12.50`

Saida:

- `a quantidade de carne para 6 pessoas no churrasco é: 3.00 kg`
- `O custo total da carne para 6 pessoas no churrasco é: R$ 37.50`
