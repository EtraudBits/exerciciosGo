# Estudo de Robustez em Crawler com Retornos Múltiplos (Go)

Este exercício mostra um dos pilares de robustez em Go: **retornos múltiplos**, especialmente no formato `(resultado, error)`.

No arquivo `crawlerErros.go`, a função principal de busca é:

```go
func buscarConteudo(url string) (string, error)
```

Ela retorna:

- `string`: conteúdo processado (quando há sucesso)
- `error`: erro da operação (quando algo falha)

---

## Objetivo do estudo

Entender como construir um crawler mais confiável, evitando falhas silenciosas e tratando problemas no momento certo.

---

## Conceito-chave: retornos múltiplos

Em Go, o padrão mais comum para operações de IO/rede é:

```go
valor, err := funcaoQuePodeFalhar()
if err != nil {
	// tratar falha imediatamente
}
```

Isso melhora a robustez porque:

- força o tratamento explícito de erros;
- evita que o programa continue com dados inválidos;
- torna o fluxo de execução previsível.

---

## Como o código atual aplica robustez

### 1) Validação de entrada

Na função `buscarConteudo`, a URL vazia é rejeitada:

```go
if url == "" {
	return "", errors.New("URL inválida")
}
```

Resultado: evita processamento inválido já no início.

### 2) Contrato claro de sucesso/falha

- **Sucesso**: retorna `conteudo, nil`
- **Falha**: retorna valor zero + erro (`"", err`)

Esse padrão simplifica o consumo da função em `main`.

### 3) Fail fast no ponto de chamada

Em `main`, o erro é tratado logo após a chamada:

```go
texto, err := buscarConteudo("http://exemplo.com")
if err != nil {
	fmt.Printf("Falha na arquitetura: %v\n", err)
	return
}
fmt.Println("processamento:", texto)
```

Resultado: o programa encerra com segurança quando ocorre falha, sem seguir para etapas inconsistentes.

---

## Leitura arquitetural do padrão

Para um crawler real, pensar em camadas ajuda:

1. **Entrada** (URL, parâmetros)
2. **Busca** (requisição HTTP)
3. **Processamento** (parse de HTML)
4. **Persistência/saída**

Cada camada deve devolver `(dados, error)` e o chamador deve decidir:

- tentar novamente,
- registrar log,
- ignorar item,
- ou abortar execução.

---

## Limitações deste exemplo (didático)

O código é propositalmente simples e ainda não cobre:

- timeout de requisição;
- classificação de erros (rede, status HTTP, parse);
- estratégia de retry com backoff;
- logs estruturados para auditoria.

---

## Evolução recomendada para robustez de crawler

### 1) Melhorar assinatura e contexto de erro

Usar `fmt.Errorf("...: %w", err)` para encadear causa raiz.

### 2) Aplicar timeout e cancelamento

Usar `context.Context` e `http.Client{Timeout: ...}`.

### 3) Tipar categorias de erro

Permite decisão automática: reprocessar erro transitório, descartar erro definitivo.

### 4) Isolar responsabilidades

Separar funções de:

- buscar (`fetch`)
- extrair (`parse`)
- salvar (`store`)

Todas com retorno múltiplo de erro.

---

## Resumo do aprendizado

Retornos múltiplos em Go não são apenas sintaxe: são uma estratégia de engenharia para aumentar a robustez de crawlers.

A ideia principal é:

- **cada função declara claramente se deu certo ou falhou**;
- **o erro é tratado imediatamente**;
- **o sistema evita estado inconsistente**.

Esse padrão é base para escalar de um exemplo simples para crawlers de produção.
