package main // Pacote principal — necessário para executar o programa com "go run"

import "fmt" // Importa o pacote fmt, usado para imprimir no terminal

// CalculateYears calcula a idade equivalente de gatos e cachorros
// com base na idade humana fornecida (em anos).
//
// A função retorna um array com 3 posições:
// [anos humanos, anos de gato, anos de cachorro]
func CalculateYears(years int) (result [3]int) {

	// Se for o primeiro ano, retornamos diretamente os valores:
	// 1 ano humano = 15 anos de gato e 15 de cachorro
	if years == 1 {
		return [3]int{1, 15, 15}
	}

	// Se for o segundo ano:
	// somamos +9 a cada um (ou seja, 15 + 9 = 24)
	if years == 2 {
		return [3]int{2, 24, 24}
	}

	// A partir do terceiro ano:
	// - Gato ganha +4 anos por ano adicional
	// - Cachorro ganha +5 anos por ano adicional
	// Primeiro, calculamos esses valores usando a fórmula:
	// 24 + (anos - 2) * incremento
	catYears := 24 + (years-2)*4
	dogYears := 24 + (years-2)*5

	// Armazenamos os resultados em um array (fixo com 3 posições)
	// onde:
	// [0] = anos humanos
	// [1] = anos de gato
	// [2] = anos de cachorro
	result = [3]int{years, catYears, dogYears}

	// Retornamos o resultado (não precisamos escrever 'return result'
	// pois ele já está nomeado na declaração da função)
	return
}

func main() {
	// Vamos testar a função com vários exemplos para entender o comportamento:

	// Exemplo 1: 1 ano humano
	fmt.Println(CalculateYears(1)) // Esperado: [1 15 15]

	// Exemplo 2: 2 anos humanos
	fmt.Println(CalculateYears(2)) // Esperado: [2 24 24]

	// Exemplo 3: 3 anos humanos
	// Gato: 24 + (3 - 2)*4 = 28
	// Cachorro: 24 + (3 - 2)*5 = 29
	fmt.Println(CalculateYears(3)) // Esperado: [3 28 29]

	// Exemplo 4: 5 anos humanos
	// Gato: 24 + (5 - 2)*4 = 36
	// Cachorro: 24 + (5 - 2)*5 = 39
	fmt.Println(CalculateYears(5)) // Esperado: [5 36 39]

	// Exemplo 5: 10 anos humanos
	fmt.Println(CalculateYears(10)) // Esperado: [10 24+(8*4)=56, 24+(8*5)=64]
}
