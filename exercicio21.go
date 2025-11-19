package main

import "fmt"

func media(num float64) string {
	if num < 60.00 {
		return "Reprovado"
	}
	return "Aprovado"
}

func main() {

	var mediaGeral float64

/*nota1 := 60.00
nota2 := 72.50
nota3 := 35.50
nota4 := 53.50

mediaGeral = (nota1+nota2+nota3+nota4)/4*/

notas := [60.00,72.50,35.50,53.50]

for i:=0; i < 5; i++ {
	notas += i
}

mediaGeral = notas/4

fmt.Printf("A sua média foi %.2f, então você foi %s.\n", mediaGeral, media(mediaGeral))
}