package main

import "fmt"


func main () {

	var pessoa int

	fmt.Print("Quantas pessoas vao ao churrasco?")
	fmt.Scan(&pessoa) // ponteiro para ler a quantidade de pessoas (aponta para a variavel pessoa)

	const mediaPorPessoa = 0.5 // constante de 0.5 kg por pessoa tipo float64

	quantidadeCarne := float64(pessoa) * mediaPorPessoa // converte a variavel pessoa para float64 para fazer a multiplicação

	fmt.Printf("a quantidade de carne para %d pessoas no churrasco é: %.2f kg\n", pessoa, quantidadeCarne)


}