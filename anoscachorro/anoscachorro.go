package main

import (
	"fmt"
)

func main () {

	var nomeCao string
	var idade int
	
	fmt.Print("Qual o nome do cao?")
	fmt.Scan(&nomeCao)

	fmt.Print("Qual a idade dele (Anos Humano)?")
	fmt.Scan(&idade)

	anosCao := idade * 7

	fmt.Printf("A idade do cao %s, é de %d anos de cao\n", nomeCao, anosCao)
}