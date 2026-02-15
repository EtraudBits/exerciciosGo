# Simulador de Gorjeta

## Objetivo

Programa em Go que calcula a gorjeta e o total a pagar de uma conta, com base no valor da conta e na forma de pagamento escolhida pelo usuario.

## O que foi feito

- Leitura do valor da conta pelo terminal.
- Selecao da forma de pagamento por opcao numerica:
  - 1 = dinheiro
  - 2 = pix
  - 3 = outros
- Definicao automatica da porcentagem de gorjeta:
  - dinheiro/pix: 15%
  - outros: 10%
- Calculo do desconto para dinheiro/pix (10%).
- Impressao do valor da gorjeta, desconto e total final.

## O que foi usado

- `bufio.NewReader(os.Stdin)` para leitura de entrada do usuario.
- `strings.TrimSpace` para limpar a entrada.
- `reaisParaCentavos` para converter texto em centavos (`int`) e evitar erros de ponto flutuante.
- Funcoes auxiliares:
  - `calcularGorgeta(totalContaCentavos, porcentagem)`
  - `formaPagamento(forma, totalContaCentavos)`

## Regras de negocio

- Gorjeta:
  - 15% quando a forma de pagamento for dinheiro ou pix.
  - 10% quando a forma de pagamento for outros.
- Desconto:
  - 10% apenas para dinheiro ou pix.

## Fluxo do programa

1. Solicita o valor da conta.
2. Solicita a forma de pagamento por opcao (1, 2 ou 3).
3. Converte o valor da conta para centavos.
4. Calcula a gorjeta conforme a forma de pagamento.
5. Calcula o desconto (se aplicavel).
6. Mostra os valores finais no terminal.

## Como executar

No diretorio do programa:

```bash
go run simuladorGorgeta.go
```

## Exemplo de uso

Entrada:

- Valor da conta: 100
- Forma de pagamento: 1

Saida esperada:

- Gorjeta: 15.00
- Desconto: 10.00
- Total: 105.00

## Observacoes

- O programa valida entradas basicas e repete a pergunta em caso de erro.
- Os calculos sao feitos em centavos e convertidos para reais apenas na exibicao.
