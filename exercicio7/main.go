package main // Pacote principal — necessário para poder executar o programa com "go run"

import "fmt" // Importa o pacote fmt para exibir resultados no console

// A função Invert recebe uma lista (slice) de números inteiros
// e devolve uma nova lista com todos os sinais invertidos.
func Invert(numeros []int) []int {
	// Cria uma nova lista vazia que guardará os números invertidos.
	// Ela começa vazia, mas ao longo do código será preenchida com append().
	var numInvertidos []int

	// Usa o loop "for range" para percorrer todos os elementos da lista "numeros".
	// O "_" significa que ignoramos o índice (posição) e usamos apenas o valor.
	for _, numero := range numeros {
		// Multiplica cada número por -1 para inverter o sinal.
		// Exemplo: 5 * (-1) = -5   e   -3 * (-1) = 3
		numInvertido := numero * -1

		// Adiciona o número invertido à lista "numInvertidos".
		// append() serve para adicionar elementos ao final de um slice.
		numInvertidos = append(numInvertidos, numInvertido)
	}

	// Retorna a nova lista com todos os valores invertidos.
	return numInvertidos
}



func main() {
	// Exemplo 1: lista com números positivos e negativos
	numeros1 := []int{1, -2, 3, -4, 5}
	fmt.Println("Lista original:", numeros1)
	fmt.Println("Lista invertida:", Invert(numeros1)) // chama a função Invert 
	// Esperado: [-1 2 -3 4 -5]

	// Exemplo 2: todos positivos
	numeros2 := []int{10, 20, 30}
	fmt.Println("\nLista original:", numeros2)
	fmt.Println("Lista invertida:", Invert(numeros2))
	// Esperado: [-10 -20 -30]

	// Exemplo 3: todos negativos
	numeros3 := []int{-7, -8, -9}
	fmt.Println("\nLista original:", numeros3)
	fmt.Println("Lista invertida:", Invert(numeros3))
	// Esperado: [7 8 9]

	// Exemplo 4: lista vazia (caso limite)
	numeros4 := []int{}
	fmt.Println("\nLista original:", numeros4)
	fmt.Println("Lista invertida:", Invert(numeros4))
	// Esperado: []
}
