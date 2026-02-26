package main

import "fmt"

// explorarSite simula a nevegação resursica de um crawler
func explorarSite(url string, profundidade int) {
	// 1. Caso base: (entenharia de segurança)
	// se não pararmos, a stack frame crescerá até o programa quebrar
	if profundidade == 0 {
		fmt.Printf("-> Limite atingido em: %s\n", url)
		return
	}

	// 2. Ação da Função: simula a exploração do site
	fmt.Printf("Explorando: %s (Nível: %d)\n", url, profundidade)

	// 3. Passo Recursivo
	//criamos um novo link e chamamos a função novamente.
	// Note que reduzimos a profundidade: isso garante que chegue ao caso base.
	proximoLink := url + "/Link"
	explorarSite(proximoLink, profundidade-1)

	// 4. Ponto de retorno
	//Esta lina só executa QUANDO a chamada acima terminar, ou seja, quando chegar ao caso base e começar a retornar.
	fmt.Printf("Finalizada a exploração de: %s\n", url)
}

func main() {
	explorarSite("http://google.com", 3)
}