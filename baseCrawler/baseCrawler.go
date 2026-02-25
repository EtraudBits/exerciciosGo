package main

import "fmt"

// percorrerLinks é o coração do crawler.
// Ela usa recursao para simular a navegação em profunidade.
func percorrerLinks(url string, profundidade int) {
	// Passo 1: Condição de Parada (Engenharia de segurança)
	// sem isso, a stack (Pilha) transbora e o progrma trava (stack overflow)
	if profundidade <= 0 {
		return
	}
	fmt.Printf("Visitando: %s (Nível: %d)\n", url, profundidade)

	//simulção: extraimos um novo link da página atual
	novoLink := url + "/sub-pagina"

	// Passo 2: A chamada recursiva
	// A função chama a si mesma para "entrar" no novo link.
	// Um novo stack Frame é criado para cada nivel de profundidade.
	percorrerLinks(novoLink, profundidade-1)
	//Passo 3: o Retorno
	// quando a profundidade chega a 0, os frames são destruídos um por um.
}

func main() {
	percorrerLinks ("http://aprendago.com", 10)
}