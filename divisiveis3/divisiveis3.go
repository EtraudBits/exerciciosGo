// Você e seus colegas querem criar um codigo que exiba todos os números entre 1 e 100 que sejam divisíveis por 3.
package main

import "fmt"


func main() {

	// Para isso, você decide usar um loop for para iterar de 1 a 100 e um operador de módulo para verificar se o número é divisível por 3.
	for i := 1; i <= 100; i++ {
		if i%3 == 0 { // O operador de módulo (%) retorna o resto da divisão de i por 3. Se o resultado for 0, significa que i é divisível por 3.
			fmt.Println(i) // Se o número for divisível por 3, ele é impresso na tela usando fmt.Println.
		}
	}

}