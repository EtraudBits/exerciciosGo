package main // Declara que este arquivo pertece ao pacote "main", ponto de entrada de um programa Go.

import (
	"fmt" //Importa o pacote "fmt" para imprimir na tela e formatar strings.
)

// Maps é uma funão que recebe um slice [] e retorna um novo slice []
// com cada elemento dobrado
// assinatura: recebe []int (lista de inteiros) e retorna []int (lista de inteiros)
func Maps(number []int) []int{ 
	//cria um slice vazio de inteiros para armazenar os resultados.
	//vamos preencher-lo com os numeros dobrados
	dobro := []int{}

// percorre o slice "numbers" usando range.
//"_" ignora o indice; numbers é o valor atual.
for _, numbers := range number {
	// calcula o dobro do valor atual (numbers * 2) e adiciona ao slice "dobro"
	dobro = append(dobro, numbers * 2)
}
//retorna o slice preenchido com todos os valores dobrados
return dobro
}

func main() {
	// Declara e inicializa um slice de inteiros com alguns valores de exemplo.
	numbers := []int{1, 2, 3, 4, 5}

	//chama a função Maps passando "numbers"
	//o retorno (slice com valores dobrados) é armazenado em "dobros".
	dobros := Maps(numbers)

	//imprime os valores originais e os valores dobrados na tela.
	fmt.Println("Numero:", numbers, "- Dobro:", dobros)
}