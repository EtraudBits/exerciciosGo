package main

import (
	"errors" // usado para sinalizar falhas de engenharia
	"fmt"
)

// calcularVolume cilindro: isolamos a matema´tica daexibição (arquetetura)
func calcularVolume(raio, altura float64) (float64, error) {
	// Passo 1: Validadação de seguranaça (SAnity check)
	if raio <= 0 || altura <= 0 {
		return 0, errors.New("dimensões devem ser maiores que zero")
	}
	// Passo 2: Processamento
	const pi = 3.1459
	volume := pi * (raio*raio) * altura

	// Passo 3: retorno do contrato (sucesso)
	return volume, nil
}

func main() {
	// Chamada da função: a 'main' não sabe COMO  o volume é calculado,
	// ela apenas recebe o resultado. Isso é Encapsulamento.

	vol, err := calcularVolume(5.0, 10.0) // 5.0 é o raio, 10.0 é a altura
	// tratmento de Erro (Padrao de ouro do GO)
	if err != nil {
		fmt.Println("Falha no Sistema:", err)
		return
	}
	fmt.Printf("Volume Calculado: %.2f m³\n", vol)
}



