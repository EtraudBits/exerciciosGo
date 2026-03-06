# Tratamento de erros em Go (Estratégias na prática)

Este exercício foi criado para **ensinar como tratar erros de forma intencional** em Go, usando um cenário simples de rede (requisição HTTP).

O foco aqui não é “evitar erro”, e sim **saber reagir ao erro com clareza**.

---

## Objetivo do exercício

Demonstrar estratégias de tratamento de erros que aparecem no dia a dia:

1. **Programar para o erro** (antecipar falhas)
2. **Repetir** (retry, quando fizer sentido)
3. **Encerrar** (falha crítica)
4. **Logar e continuar/parar fluxo local**
5. **Ignorar (raro)**

No código atual, as estratégias **mais visíveis** são:

- Estratégia 1: programar para o erro com contexto
- Estratégia 4: logar erro e interromper o fluxo da execução atual

---

## O que foi feito no código

### 1) Função dedicada para acesso HTTP

Foi criada a função `buscarHTML(url string) (*http.Response, error)`.

Ela:

- recebe uma URL
- executa `http.Get(url)`
- retorna dois valores: resposta HTTP e erro

Esse padrão (`valor, erro`) é o padrão idiomático de Go.

### 2) Tratamento com enriquecimento de contexto

Ao invés de retornar somente o erro original, o código faz:

`fmt.Errorf("falha na rede ao acessar %s: %v", url, err)`

Isso melhora muito o diagnóstico, porque informa:

- **o que falhou** (acesso de rede)
- **onde falhou** (na URL específica)
- **qual foi o erro original**

### 3) Simulação de erro realista

No `main`, a URL usada é propositalmente inválida (`http://site-inexistente.com`) para forçar uma falha e observar o comportamento.

### 4) Log de engenharia no ponto de decisão

Quando `buscarHTML` retorna erro, o `main` faz:

- log com `fmt.Printf("LOG DE ENGENHARIA: ...")`
- `return` para encerrar aquele fluxo com segurança

Isso evita continuar a execução com estado inválido.

---

## Estratégias de erro: conceito aplicado

### Estratégia 1 — Programar

**Aplicada no código.**

Como foi aplicada:

- validação implícita via retorno de erro do `http.Get`
- retorno de erro com mensagem contextualizada

Quando usar:

- sempre que uma operação externa pode falhar (rede, disco, banco, API)

### Estratégia 2 — Repetir

**Não implementada neste exemplo**, mas importante.

Seria usada quando falhas são temporárias (timeout, indisponibilidade momentânea).

Exemplo conceitual:

- tentar até `N` vezes
- aplicar intervalo crescente entre tentativas (backoff)

### Estratégia 3 — Encerrar

**Parcialmente aplicada** com `return` no `main`.

Em aplicações maiores, poderia ser encerramento do processo quando o erro é crítico e impede funcionamento correto.

### Estratégia 4 — Logar e continuar/parar fluxo local

**Aplicada no código.**

Foi feito log com contexto para facilitar troubleshooting e a execução atual é interrompida de forma controlada.

### Estratégia 5 — Ignorar (raro)

**Não aplicada (corretamente evitada).**

Ignorar erro é perigoso e só deve acontecer quando:

- o erro é realmente irrelevante para o objetivo
- e está documentado por que foi ignorado

---

## Pacotes e recursos usados

- `net/http`: para simular chamada HTTP real (`http.Get`)
- `fmt`: para formatação de mensagens e logs
- retorno múltiplo de Go: `(*http.Response, error)`
- construção de erro contextual com `fmt.Errorf`

---

## Fluxo de execução (passo a passo)

1. `main` define a URL.
2. `main` chama `buscarHTML(url)`.
3. `buscarHTML` faz `http.Get(url)`.
4. Se houver erro:
   - cria mensagem contextual
   - retorna `nil, erro`
5. `main` recebe o erro, registra log e encerra o fluxo.
6. Se não houver erro, imprime sucesso.

---

## Resultado esperado ao executar

Como a URL é inválida, o comportamento esperado é:

- mensagem de log indicando falha de rede
- término controlado da execução

Comando:

`go run errosEstrategias1.go`

---

## Aprendizado principal

Este exercício ensina que erro em Go é parte do fluxo normal do programa.

A prática correta é:

- detectar cedo
- adicionar contexto útil
- decidir estratégia adequada (repetir, encerrar, logar, etc.)
- evitar ignorar erros sem justificativa

Isso gera código mais confiável, mais fácil de manter e mais fácil de depurar.
