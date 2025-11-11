package main

import "fmt"

func sinalContrario (numero int) int {
	if numero > 0 {
		return numero * (-1)
	}
	return numero
}

func main () {

	numero := 1
	fmt.Println (sinalContrario(numero))
}