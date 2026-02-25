# Estudo de Go: Funções com exemplo de Web Crawler

## O que estamos estudando

Neste exercício, o foco principal é **Funções** em Go.

Usamos um exemplo simples de **Web Crawler** para praticar:

- declaração de função;
- parâmetros;
- chamada de função;
- condição de parada;
- recursão (função chamando a si mesma).

---

## Código de referência

```go
func percorrerLinks(url string, profundidade int) {
	if profundidade <= 0 {
		return
	}

	fmt.Printf("Visitando: %s (Nível: %d)\n", url, profundidade)

	novoLink := url + "sub-pagina"

	percorrerLinks(novoLink, profundidade-1)
}
```

---

## Entendendo a função `percorrerLinks`

### 1) Nome da função

`percorrerLinks` descreve a responsabilidade da função: visitar links em sequência.

### 2) Parâmetros

```go
url string, profundidade int
```

- `url`: endereço atual sendo visitado;
- `profundidade`: quantos níveis ainda podemos descer.

### 3) Retorno

A função não retorna valor explícito (retorno vazio), apenas executa ações:

- imprime informações no terminal;
- chama a próxima etapa da navegação.

### 4) Condição de parada

```go
if profundidade <= 0 {
	return
}
```

Essa condição evita recursão infinita. Quando a profundidade chega a 0, a execução volta para as chamadas anteriores.

### 5) Chamada recursiva

```go
percorrerLinks(novoLink, profundidade-1)
```

Aqui está o ponto-chave:

- a função chama ela mesma;
- `profundidade-1` reduz 1 nível a cada chamada;
- isso garante que a condição de parada será alcançada.

---

## Fluxo de execução com `profundidade = 3`

Chamada inicial:

```go
percorrerLinks("http://aprendago.com", 3)
```

Ordem prática:

1. Nível 3: imprime e chama nível 2;
2. Nível 2: imprime e chama nível 1;
3. Nível 1: imprime e chama nível 0;
4. Nível 0: para (`return`) e começa o retorno da pilha de chamadas.

Resultado: 3 visitas impressas.

---

## O que este exemplo ensina sobre Funções

- funções organizam o problema em blocos reutilizáveis;
- parâmetros permitem controlar o comportamento da função;
- funções podem trabalhar em conjunto com lógica de repetição via recursão;
- toda recursão precisa de uma condição de parada segura.

---

## Limitação deste crawler (didático)

Este exemplo é uma simulação para aprendizado. Ele **não busca links reais da internet**; apenas monta novas strings com `"sub-pagina"`.

---

## Como executar

Dentro da pasta `baseCrawler`:

```bash
go run baseCrawler.go
```

---

## Sugestões para praticar mais

1. Trocar a profundidade para `1`, `5` e `10` e observar o resultado.
2. Adicionar um contador de visitas.
3. Mudar o texto de saída para mostrar também o link pai.
4. Criar uma versão com laço (`for`) e comparar com recursão.
