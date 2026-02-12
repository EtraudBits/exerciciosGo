package main

import "fmt"


func main () {

	var pessoa int
	var precoCarne float64

	fmt.Print("Quantas pessoas vao ao churrasco?")
	fmt.Scan(&pessoa) // ponteiro para ler a quantidade de pessoas (aponta para a variavel pessoa)

	fmt.Print("Qual o preço da Carne (kg)?")
	fmt.Scan(&precoCarne) // ponteiro para ler o preço da carne (aponta para a variavel precoCarne)

	const mediaPorPessoa = 0.5 // constante de 0.5 kg por pessoa tipo float64

	quantidadeCarne := float64(pessoa) * mediaPorPessoa // converte a variavel pessoa para float64 para fazer a multiplicação

	custoTotalDaCarne := quantidadeCarne * precoCarne // calcula o custo total da carne

	fmt.Printf("a quantidade de carne para %d pessoas no churrasco é: %.2f kg\n", pessoa, quantidadeCarne)

	fmt.Printf("O custo total da carne para %d pessoas no churrasco é: R$ %.2f\n", pessoa, custoTotalDaCarne)


}