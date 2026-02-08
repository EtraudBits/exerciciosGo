package main // Pacote principal, necessário para programas executáveis

import (
    "fmt"      // Pacote para imprimir no console
    "strconv"  // Pacote usado para conversão entre string e tipos numéricos
)

// StringToNumber converte uma string que contém um número em um tipo inteiro.
func StringToNumber(str string) int {
    // strconv.Atoi tenta converter a string para um número inteiro.
    // O nome Atoi vem de "ASCII to integer".
    //
    // Ele retorna dois valores:
    // - O número convertido (se a conversão for bem-sucedida)
    // - Um erro (se a string não puder ser convertida)
    //
    // Como estamos confiando que a string é válida, ignoramos o erro usando "_".
    number, _ := strconv.Atoi(str)

    // Retorna o número convertido
    return number
}

func main() {
    // Declara uma string representando um número
    texto := "42"

    // Chama a função para converter a string em número
    resultado := StringToNumber(texto)

    // Mostra os valores antes e depois da conversão
    fmt.Println("String original:", texto)
    fmt.Println("Número convertido:", resultado)

    // Outros exemplos:
    fmt.Println("StringToNumber(\"0\") =", StringToNumber("0"))   // Esperado: 0
    fmt.Println("StringToNumber(\"-15\") =", StringToNumber("-15")) // Esperado: -15
}
