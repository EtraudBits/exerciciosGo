package main

import (
	"fmt"
)

func removerExtremos(palavra string) string {
	if len(palavra) <= 2 { //len(palavra) retorna o tamanho da string
		return ""
	}
	return palavra[1 : len(palavra) -1] // começa do indice 1 e vai até o penultimo indice (ou seja, exclui o ultimo inidice)
}

func main() {

	fmt.Println(removerExtremos("Duarte"))
	fmt.Println(removerExtremos("Iris"))
	fmt.Println(removerExtremos("OK"))
}