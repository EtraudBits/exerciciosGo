# Estrategias de Erros em Go (Guia de Estudo)

Este material resume, de forma didatica, as estrategias de tratamento de erros aplicadas nos exercicios deste repositorio.

O objetivo nao e "eliminar erros", e sim tratar falhas com clareza, previsibilidade e seguranca.

## Visao Geral

Nos exercicios, aparecem as 5 estrategias classicas:

1. Programar para o erro
2. Repetir (retry)
3. Encerrar execucao
4. Logar e continuar/parar fluxo local
5. Ignorar (raro)

Arquivos analisados:

- `erros1/errosEstrategias1.go`
- `erros2/errosEstrategias2.go`
- `crawlerErrors/crawlerErros.go`

## O que foi usado e como

### 1) Programar para o erro

Onde aparece:

- `erros1/errosEstrategias1.go`
- `crawlerErrors/crawlerErros.go`

Como foi aplicado:

- Funcoes retornam `(..., error)` para explicitar sucesso/falha.
- Entradas invalidas sao tratadas cedo (ex.: URL vazia).
- Erros recebem contexto para facilitar diagnostico.

Exemplo real (`erros1`):

```go
resp, err := http.Get(url)
if err != nil {
	return nil, fmt.Errorf("falha na rede ao acessar %s: %v", url, err)
}
```

Exemplo real (`crawlerErrors`):

```go
if url == "" {
	return "", errors.New("URL invalida")
}
```

Aprendizado:

- O erro vira parte normal do fluxo de controle.
- O chamador decide o que fazer, em vez de "esconder" a falha.

### 2) Repetir (Retry)

Onde aparece:

- `erros2/errosEstrategias2.go`

Como foi aplicado:

- A funcao `buscarComRetry` tenta varias vezes dentro de uma janela de tempo (`1 minuto`).
- A cada falha, registra log e espera com backoff exponencial (`1s`, `2s`, `4s`...).

Trecho real:

```go
for tentativas := 0; time.Now().Before(deadline); tentativas++ {
	resp, err := http.Get(url)
	if err == nil {
		return resp, nil
	}
	log.Printf("Servidor nao respondeu (%s), tentando novamente...", err)
	time.Sleep(time.Second << tentativas)
}
```

Aprendizado:

- Retry e util para falhas temporarias (rede instavel, timeout intermitente).
- Sem limite de tempo/tentativas, retry pode travar o sistema.

### 3) Encerrar execucao

Onde aparece:

- `erros2/errosEstrategias2.go`

Como foi aplicado:

- Apos esgotar as retentativas, o erro sobe para `main`.
- Em caso critico, `log.Fatalf` registra e encerra o programa.

Trecho real:

```go
if err != nil {
	log.Fatalf("Erro critico no crowler: %v", err)
}
```

Aprendizado:

- Encerrar e correto quando nao existe caminho seguro para continuar.

### 4) Logar e parar fluxo local

Onde aparece:

- `erros1/errosEstrategias1.go`
- `crawlerErrors/crawlerErros.go`

Como foi aplicado:

- O erro e registrado no ponto de decisao.
- Em seguida, o fluxo atual e interrompido com `return`.

Trecho real (`erros1`):

```go
if err != nil {
	fmt.Printf("LOG DE ENGENHARIA: %v\n", err)
	return
}
```

Aprendizado:

- Evita continuar com dados invalidos.
- Facilita suporte e depuracao.

### 5) Ignorar erro (raro)

Onde aparece:

- Nao foi aplicado de forma explicita nos tres arquivos (o que e positivo).

Quando considerar:

- Apenas quando o erro nao impacta o objetivo principal.
- Deve ficar documentado no codigo por que foi ignorado.

## Fluxo mental recomendado para novos exercicios

1. Operacao pode falhar?
2. Se sim, retorne `error` com contexto.
3. No chamador, trate imediatamente com `if err != nil`.
4. Escolha estrategia:

- retry (falha temporaria)
- log + return (fluxo local)
- fatal/encerrar (falha critica)

5. Nunca ignore erro sem justificativa.

## Resumo pratico

- `erros1` ensina contexto de erro + log no ponto certo.
- `erros2` ensina retry com backoff + falha critica com encerramento.
- `crawlerErrors` reforca validacao de entrada e contrato `resultado, error`.

Esses tres exercicios formam uma base solida para evoluir para aplicacoes reais com tratamento de erros consistente.
