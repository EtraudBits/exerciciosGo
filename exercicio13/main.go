package main

import (
	"fmt"
	"strings"
)

// PartList recebe um slice de strings e devolve uma string que contém
// todas as formas de dividir a lista em duas partes não vazias, no formato:
// "(left, right) (left2, right2) ..."
// Ex.: ["I", "love", "coding"] -> "(I, love coding) (I love, coding)"
func PartList(listas []string) string {
	// Se a lista tem menos de 2 elementos, não existe divisão válida.
	if len(listas) < 2 {
		return ""
	}

	// finalLista acumula todas as tuplas geradas.
	var finalLista string

	// Percorremos as posições possíveis de corte da lista.
	// i começa em 1 porque a parte esquerda não pode ser vazia,
	// e vai até len(listas)-1 porque a parte direita também não pode ser vazia.
	for i := 1; i < len(listas); i++ {
		// strings.Join(listas[:i], " ") une os elementos da parte esquerda com espaço.
		left := strings.Join(listas[:i], " ")

		// strings.Join(listas[i:], " ") une os elementos da parte direita com espaço.
		right := strings.Join(listas[i:], " ")

		// fmt.Sprintf formata a tupla no formato solicitado "(left, right)".
		lista := fmt.Sprintf("(%s, %s)", left, right)

		// Para melhorar a leitura, adicionamos um espaço antes da tupla
		// a partir da segunda iteração (para separar "(...), (...)").
		if i > 1 {
			finalLista += " "
		}

		// Concatena a tupla ao resultado final.
		finalLista += lista
	}

	// Retorna a string com todas as divisões formatadas.
	return finalLista
}

func main() {
	// Testes práticos:

	ex1 := []string{"I", "love", "coding"}
	// Divisões:
	// i=1 -> left="I"         right="love coding"   => "(I, love coding)"
	// i=2 -> left="I love"    right="coding"        => "(I love, coding)"
	fmt.Println("Exemplo 1:", PartList(ex1))
	// Esperado: "(I, love coding) (I love, coding)"

	ex2 := []string{"cd", "ex", "ab", "xy"}
	// i=1 -> "(cd, ex ab xy)"
	// i=2 -> "(cd ex, ab xy)"
	// i=3 -> "(cd ex ab, xy)"
	fmt.Println("Exemplo 2:", PartList(ex2))
	// Esperado: "(cd, ex ab xy) (cd ex, ab xy) (cd ex ab, xy)"

	ex3 := []string{"Hello", "World"}
	// Apenas uma divisão possível: "(Hello, World)"
	fmt.Println("Exemplo 3:", PartList(ex3))
	// Esperado: "(Hello, World)"

	ex4 := []string{"Solo"}
	// Não há como dividir um único elemento em duas partes não vazias
	fmt.Println("Exemplo 4:", PartList(ex4))
	// Esperado: "" (string vazia)
}
