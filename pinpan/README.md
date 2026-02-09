# Pin Pan (1 a 100)

Este exercicio implementa a brincadeira em que:

- Multiplos de 3 viram "Pin"
- Multiplos de 5 viram "Pan"
- Multiplos de 3 e 5 viram "PinPan"

Os numeros que nao se encaixam nessas regras sao impressos normalmente.

## Objetivo

- Treinar loop `for`.
- Usar condicionais (`if / else if / else`).
- Aplicar o operador de modulo (`%`) para verificar multiplos.

## Como executar

No terminal, dentro da pasta `pinpan`:

```bash
go run pinpan.go
```

Opcionalmente:

```bash
go run .
```

## Saida esperada (exemplo)

```text
1
2
Pin
4
Pan
Pin
7
8
Pin
Pan
...
PinPan
```

## Estrutura

- `pinpan.go`: codigo principal do exercicio.
