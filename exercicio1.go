package main // Define o pacote principal; imprescindível para um executável em Go

import "fmt" // Importa o pacote fmt, usado para imprimir no terminal

// NumNegativo recebe um inteiro x e garante que o valor retornado seja negativo (ou zero).
func NumNegativo(x int) int {
	// Verifica se x já é menor ou igual a zero.
	if x <= 0 {
		// Se for, retorna x sem alteração (já é negativo ou zero).
		return x
	}
	// Se x for positivo, retorna o seu negativo.
	return -x
}

func main() {
	// Declara e inicializa a variável 'a' com valor 5.
	a := 5
	// Declara e inicializa a variável 'b' com valor -3.
	b := -3
	// Declara e inicializa a variável 'c' com valor 0.
	c := 0

	// Chama MakeNegative com 'a' e imprime o resultado. Aqui esperamos -5.
	fmt.Println("NumNegativo(", a, ") ->", NumNegativo(a))
	// Chama MakeNegative com 'b' e imprime o resultado. Aqui esperamos -3 (permanece -3).
	fmt.Println("NumNegativo(", b, ") ->", NumNegativo(b))
	// Chama MakeNegative com 'c' e imprime o resultado. Aqui esperamos 0.
	fmt.Println("NumNegativo(", c, ") ->", NumNegativo(c))
}