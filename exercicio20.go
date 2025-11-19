package main

import "fmt"

func move(posicaoInicial, dado int) int { //move calcula a nova posição do horoi no tabuleiro.

	avanco := dado * 2 // o heroi anda duas vezes o numero que saiu no dado

	novaPosicao := posicaoInicial + avanco // a nova posição é a posição inicial + o avanço

	return novaPosicao // retorna o resultado

}

func main ()  {

	posicao := 3 // varialvel da posição que o herois estar atualmente
	dado := 6 // valor que o dado caiu

	fmt.Println(move(posicao, dado)) // a função move dobra o valor do dado, onde a resposta será a posição inicial mais o dobro do valor do dado
	
}