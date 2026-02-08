package main // Pacote principal — necessário para executar o programa com "go run"

import "fmt" // Importa o pacote fmt, usado para imprimir resultados no console

// EvenOrOdd recebe um número inteiro e retorna se ele é "Even" (par) ou "Odd" (ímpar).
func EvenOrOdd(number int) string {

	// A operação "number % 2" calcula o resto da divisão de number por 2.
	// - Se o resto for 0 → o número é par (Even)
	// - Se o resto for diferente de 0 → o número é ímpar (Odd)
	if number%2 == 0 {
		// Caso o resto da divisão por 2 seja zero, retornamos "Even"
		return "Even"
	} else {
		// Caso contrário, o número é ímpar, retornamos "Odd"
		// Observação: como só existem duas possibilidades (par ou ímpar),
		// o "else" é opcional — poderíamos simplesmente usar outro "return" fora do if.
		return "Odd"
	}
}

func main() {
	// Aqui testamos a função com vários valores para entender o comportamento.

	// Exemplo 1: número par
	num1 := 4
	fmt.Println("O número", num1, "é:", EvenOrOdd(num1))
	// number%2 = 4%2 = 0 → retorna "Even"

	// Exemplo 2: número ímpar
	num2 := 7
	fmt.Println("O número", num2, "é:", EvenOrOdd(num2))
	// number%2 = 7%2 = 1 → retorna "Odd"

	// Exemplo 3: número negativo par
	num3 := -10
	fmt.Println("O número", num3, "é:", EvenOrOdd(num3))
	// (-10)%2 = 0 → retorna "Even"

	// Exemplo 4: número negativo ímpar
	num4 := -3
	fmt.Println("O número", num4, "é:", EvenOrOdd(num4))
	// (-3)%2 = -1 → retorna "Odd"

	// Exemplo 5: zero (0 é considerado par)
	num5 := 0
	fmt.Println("O número", num5, "é:", EvenOrOdd(num5))
	// 0%2 = 0 → retorna "Even"
}
