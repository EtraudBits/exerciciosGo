package main

import "fmt"

// SumPositive soma apenas os números positivos de um slice.
func SumPositive(numbers []int) int {
    var soma int // acumulador; começa em 0

    // for-range: percorre CADA elemento do slice
    for _, v := range numbers {
        // só soma os valores maiores que zero
        if v > 0 {
            soma += v
        }
    }

    // retorna o resultado acumulado
    return soma
}

func main() {
    exemplo := []int{1, -4, 7, 12}
    fmt.Println("Exemplo:", exemplo)
    // passo a passo do exemplo:
    // soma começa 0
    // v=1  -> soma = 0 + 1 = 1
    // v=-4 -> soma = 1 (ignora)
    // v=7  -> soma = 1 + 7 = 8
    // v=12 -> soma = 8 + 12 = 20
    fmt.Println("SumPositive(exemplo) =", SumPositive(exemplo)) // Esperado: 20

    // outros testes
    fmt.Println("SumPositive([]int{}) =", SumPositive([]int{}))         // 0 (vazio)
    fmt.Println("SumPositive([]int{-1,-2,-3}) =", SumPositive([]int{-1, -2, -3})) // 0 (todos negativos)
    fmt.Println("SumPositive([]int{5,10}) =", SumPositive([]int{5, 10})) // 15
}
