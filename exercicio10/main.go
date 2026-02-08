package main // Pacote principal — necessário para executar o código com "go run"

import "fmt" // Importa o pacote fmt, usado para imprimir resultados no terminal

// Função Rps (Rock, Paper, Scissors)
// Simula o jogo "Pedra, Papel e Tesoura"
// Recebe duas strings — p1 e p2 — que representam as jogadas dos jogadores
// Retorna uma string indicando quem ganhou (Player 1, Player 2) ou se deu empate.
func Rps(p1, p2 string) string {

	// 🟢 1º Caso: Empate
	// Se ambos escolherem a mesma opção, o resultado é um "Draw!"
	if p1 == p2 {
		return "Draw!" // Retorna "Draw!" (empate)
	}

	// 🧱✂️ Caso 1: Pedra (rock) ganha da tesoura (scissors)
	if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	}

	// ✂️📄 Caso 2: Tesoura (scissors) ganha do papel (paper)
	if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	}

	// 📄🧱 Caso 3: Papel (paper) ganha da pedra (rock)
	if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	}

	// 🔴 Se nenhum dos casos anteriores for verdadeiro,
	// significa que o Player 2 venceu.
	return "Player 2 won!"
}

func main() {
	// 🧪 Testes práticos para entender o comportamento da função:

	fmt.Println("Teste 1:", Rps("rock", "scissors"))  // Esperado: Player 1 won!
	fmt.Println("Teste 2:", Rps("scissors", "paper")) // Esperado: Player 1 won!
	fmt.Println("Teste 3:", Rps("paper", "rock"))     // Esperado: Player 1 won!
	fmt.Println("Teste 4:", Rps("rock", "paper"))     // Esperado: Player 2 won!
	fmt.Println("Teste 5:", Rps("scissors", "rock"))  // Esperado: Player 2 won!
	fmt.Println("Teste 6:", Rps("paper", "scissors")) // Esperado: Player 2 won!
	fmt.Println("Teste 7:", Rps("paper", "paper"))    // Esperado: Draw!
}
