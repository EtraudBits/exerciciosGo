package main

import "fmt"

// SortMyString recebe uma string e separa as letras que estão
// nas posições pares e ímpares, retornando-as no formato:
// "letrasPares letrasÍmpares"
// Exemplo: "CodeWars" -> "CdWr oeas"
func SortMyString(divisao string) string {
	// Variáveis para armazenar as letras pares e ímpares separadamente.
	var letraPar, letraImpar string

	// Percorre cada caractere da string com base no índice i.
	for i := 0; i < len(divisao); i++ {
		// A condição (i % 2 == 0) verifica se o índice é par.
		// Se for par, adiciona o caractere atual em letraPar.
		if i%2 == 0 {
			letraPar += string(divisao[i])
		} else {
			// Caso contrário (índice ímpar), adiciona em letraImpar.
			letraImpar += string(divisao[i])
		}
	}

	// fmt.Sprintf formata o retorno com um espaço entre os dois grupos.
	// "%s %s" insere letraPar e letraImpar, separados por espaço.
	return fmt.Sprintf("%s %s", letraPar, letraImpar)
}

func main() {
	// Testes práticos:

	ex1 := "CodeWars"
	// Índices: 0 1 2 3 4 5 6 7
	// Letras:  C o d e W a r s
	// Pares:   C d W r
	// Ímpares: o e a s
	fmt.Println("Exemplo 1:", SortMyString(ex1))
	// Esperado: "CdWr oeas"

	ex2 := "abcdef"
	// Pares: a c e
	// Ímpares: b d f
	fmt.Println("Exemplo 2:", SortMyString(ex2))
	// Esperado: "ace bdf"

	ex3 := "A"
	// Apenas uma letra (índice 0, par)
	fmt.Println("Exemplo 3:", SortMyString(ex3))
	// Esperado: "A " (sem letras ímpares, mas com espaço no retorno)

	ex4 := "GoLang"
	// Índices: 0 1 2 3 4 5
	// Letras:  G o L a n g
	// Pares: G L n
	// Ímpares: o a g
	fmt.Println("Exemplo 4:", SortMyString(ex4))
	// Esperado: "GLn oag"

	ex5 := ""
	// String vazia retorna duas strings vazias com espaço entre.
	fmt.Println("Exemplo 5:", SortMyString(ex5))
	// Esperado: " "
}
