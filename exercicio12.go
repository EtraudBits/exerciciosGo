package main // Pacote principal — necessário para executar o programa com "go run"

import (
	"fmt"     // Importa o pacote fmt, usado para imprimir no terminal
	"strings" // Importa o pacote strings, que tem funções para manipulação de texto
	"strconv" // Importa strconv, para converter texto (string) em número (int)
)

// Função Points:
// Recebe uma lista de resultados de jogos no formato ["3:1", "2:2", "0:1", ...]
// Cada string representa um placar: "nossos pontos : pontos do adversário"
//
// Regras:
// - Vitória → +3 pontos
// - Empate → +1 ponto
// - Derrota → +0 pontos
func Points(jogos []string) int {

	var pontos int // Cria uma variável para acumular o total de pontos

	// Percorre toda a lista de jogos usando "for range"
	for _, jogo := range jogos {

		// Divide a string onde há ":", resultando em uma lista com dois valores: ["3", "1"]
		placar := strings.Split(jogo, ":")

		// Converte os valores do placar (strings) para números inteiros
		nossoPlacar, _ := strconv.Atoi(placar[0])     // Primeiro valor (nosso time)
		adversarioPlacar, _ := strconv.Atoi(placar[1]) // Segundo valor (time adversário)

		// ⚽ Verificações de resultado:

		if nossoPlacar > adversarioPlacar {
			// Vitória → soma 3 pontos
			pontos += 3
		} else if nossoPlacar == adversarioPlacar {
			// Empate → soma 1 ponto
			pontos += 1
		}
		// Caso contrário (derrota), não soma nada → +0 pontos
	}

	// Retorna o total de pontos acumulados após todos os jogos
	return pontos
}

func main() {
	// 🧪 Testes práticos:

	fmt.Println("Teste 1:", Points([]string{"3:1", "2:2", "0:1"}))               // Esperado: 4 (3+1+0)
	fmt.Println("Teste 2:", Points([]string{"1:0", "2:0", "3:0"}))               // Esperado: 9 (3+3+3)
	fmt.Println("Teste 3:", Points([]string{"1:1", "2:2", "3:3"}))               // Esperado: 3 (1+1+1)
	fmt.Println("Teste 4:", Points([]string{"0:1", "0:2", "1:3"}))               // Esperado: 0 (todas derrotas)
	fmt.Println("Teste 5:", Points([]string{"1:0", "2:2", "0:1", "3:1"}))        // Esperado: 7 (3+1+0+3)
	fmt.Println("Teste 6:", Points([]string{}))                                  // Esperado: 0 (lista vazia)
}
