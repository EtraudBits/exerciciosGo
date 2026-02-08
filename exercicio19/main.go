package main

import "fmt"

func sinalContrario (numero int) int {
	/*if numero > 0 { //se o numero for maior que zero
		return numero * (-1) //retorna a multiplicação dele por -1 (um negativo)
	}
	return numero //retorna o numero negativo
}*/
return -numero //ou simplesmente retorna o numero negativo
}

func main () {

	numero := 100
	fmt.Println (sinalContrario(numero))
}