// Voce e seus colegas querem criar em formato de codigo uma antiga brincadeira: ao falar os números sempre que aparecer um numero multiplo de 3, o participante deve falar "Pin" e nos multiplos de 5 o participante deve falar "Pan". Então, seu programa de ve imprimir onumeros de 1 a 100 e nos casos cidtados ,não devem aparecer os números, mas sim o que o participante deve falar.

package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 { // Verifica se o número é múltiplo de 3 e 5
			fmt.Println("PinPan")
		} else if i%3 == 0 { // Verifica se o número é múltiplo de 3
			fmt.Println("Pin")
		} else if i%5 == 0 { // Verifica se o número é múltiplo de 5
			fmt.Println("Pan")
		} else {
			fmt.Println(i) // caso contrario os numeros que não forem multiplos de 3 ou 5 devem ser impressos normalmente sem os numeros substituidos por Pin ou Pan
			
		}
	}	
}

	
