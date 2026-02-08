package main // Pacote principal: permite executar o programa

import "fmt" // Importa o pacote fmt, usado para imprimir no terminal

// CountPositivesSumNegatives percorre um slice de inteiros e retorna:
// [quantidade de números positivos, soma de números negativos]
func CountPositivesSumNegatives(numbers []int) []int {
	// Verifica se a lista está vazia (nenhum número foi passado)
	if len(numbers) == 0 {
		// Retorna um slice vazio, pois não há o que contar nem somar
		return []int{}
	}

	// Declara variáveis para armazenar resultados intermediários
	var qntPositivos int  // vai contar quantos números são positivos (> 0)
	var somaNegativos int // vai acumular a soma dos números negativos (< 0)

	// Percorre cada número da lista usando o laço for range
	for _, number := range numbers {
		// "_" ignora o índice; "number" representa o valor atual

		// Se o número for positivo (> 0), adicionamos 1 ao contador
		if number > 0 {
			qntPositivos++
		}

		// Se o número for negativo (< 0), adicionamos o valor à soma
		if number < 0 {
			somaNegativos += number
		}
	}

	// Retorna uma lista (slice) com dois valores:
	// o primeiro é a quantidade de positivos
	// o segundo é a soma dos negativos
	return []int{qntPositivos, somaNegativos}
}

func main() {
	// Declara um exemplo de lista com números positivos e negativos
	numeros := []int{1, -2, 3, 4, -5, -6}

	// Chama a função e armazena o resultado
	resultado := CountPositivesSumNegatives(numeros)

	// Imprime a lista original
	fmt.Println("Lista:", numeros)

	// Imprime o resultado formatado
	// Esperado: [3, -13] → (3 positivos, soma dos negativos = -13)
	fmt.Println("Resultado [positivos, soma_negativos]:", resultado)
}
