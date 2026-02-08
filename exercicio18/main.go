package main

import "fmt"

func medias (notas int) rune {

	if notas <= 100 && notas >=90 {
		return 'A'
	}
	if notas < 90 && notas >= 80 {
		return 'B'
	}
	if notas < 80 && notas >= 70 {
		return 'C'
	}
	if notas < 70 && notas >= 60 {
		return 'D'
	}
	return 'F'
	}


func main () {

	a := 60
	b := 70
	c := 30

	media := (a + b + c) / 3

	fmt.Println("a média é: ", media)
	fmt.Println("você tirou: ", string (medias(media)))
}