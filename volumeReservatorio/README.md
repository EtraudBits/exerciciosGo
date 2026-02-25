# Estudo de Go: Funções com Arquitetura de Validação

## Tema do exercício

Neste código, estamos estudando **Funções** em Go usando o exemplo de cálculo do **volume de um cilindro**.

A ideia central é aplicar uma pequena **arquitetura de validação**:

- validar os dados de entrada;
- processar o cálculo somente com dados válidos;
- retornar resultado e possível erro para quem chamou a função.

---

## O que a função faz

Função principal do domínio:

```go
func calcularVolume(raio, altura float64) (float64, error)
```

Essa assinatura ensina três pontos importantes sobre funções:

1. **Parâmetros tipados**: `raio` e `altura` são `float64`.
2. **Retorno múltiplo**: retorna o volume (`float64`) e um erro (`error`).
3. **Contrato claro**: se os dados forem inválidos, retorna erro; se forem válidos, retorna volume e `nil` no erro.

---

## Arquitetura de Validação (3 passos)

### 1) Validação de entrada

```go
if raio <= 0 || altura <= 0 {
	return 0, errors.New("dimensões devem ser maiores que zero")
}
```

Antes de calcular, a função protege o sistema de entradas inválidas.

### 2) Processamento

```go
const pi = 3.1459
volume := pi * (raio * raio) * altura
```

Com os dados válidos, aplica a fórmula do cilindro:

$$V = \pi r^2 h$$

### 3) Retorno do contrato

```go
return volume, nil
```

Retorna valor calculado e informa que não houve falha.

---

## Papel da `main`

A função `main` apenas **orquestra**:

- chama `calcularVolume`;
- verifica o `err`;
- imprime sucesso ou falha.

Isso reforça um princípio importante: separar **regra de negócio** (cálculo + validação) da **interface** (entrada/saída no terminal).

---

## Fluxo de execução do exemplo

Chamada atual:

```go
vol, err := calcularVolume(5.0, 10.0)
```

Fluxo:

1. `main` chama a função.
2. A função valida raio e altura.
3. Como os valores são válidos, calcula o volume.
4. Retorna `vol` e `nil`.
5. `main` imprime: `Volume Calculado: ... m³`.

---

## Como executar

Na pasta `volumeReservatorio`:

```bash
go run volumeReservatorio.go
```

---

## Prática sugerida

1. Teste com `calcularVolume(-2, 10)` para ver o fluxo de erro.
2. Teste com `calcularVolume(3, 0)` para validar altura inválida.
3. Troque valores para observar mudanças no resultado.
4. (Opcional) Substitua `pi` por `math.Pi` para maior precisão.
