package main

import "fmt"


func main () {

	var pessoa int

	fmt.Print("Quantas pessoas vao ao churrasco?")
	fmt.Scan(&pessoa)

	const mediaPorPessoa = 0.5

	quantidadeCarne := float64(pessoa) * mediaPorPessoa

	fmt.Printf("a quantidade de carne para %d pessoas no churrasco é: %.2f kg\n", pessoa, quantidadeCarne)


}