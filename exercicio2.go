package main // Define o pacote principal — necessário para executar o programa

import "fmt" // Importa o pacote fmt, usado para imprimir no terminal

// Somacao recebe um número inteiro n e calcula a soma de todos os números de 1 até n.
func Somacao(n int) int {
	// Declara a variável 'soma' fora do laço 'for',
	// para que ela possa ser acessada depois do loop terminar.
	var soma int

	// Cria um laço que começa em 1 e vai até 'n' (inclusive).
	for i := 1; i <= n; i++ {
		// A cada repetição, adiciona o valor de 'i' à variável 'soma'.
		// É o mesmo que escrever: soma = soma + i
		soma += i
	}

	// Retorna o valor final acumulado em 'soma'.
	return soma
}

func main() {
	// Declara a variável 'n' e define o valor até onde queremos somar.
	n := 3

	// Chama a função Somacao passando 'n' e guarda o resultado na variável 'resultado'.
	resultado := Somacao(n)

	// Imprime o resultado final na tela.
	fmt.Println("A soma de 1 até", n, "é:", resultado)
}
