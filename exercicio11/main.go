package main // Pacote principal — necessário para executar o programa com "go run"

import "fmt" // Importa o pacote fmt, usado para imprimir saídas no terminal

// Função Well: analisa uma lista de ideias e decide o resultado baseado em quantas são "good".
//
// - Se houver nenhuma ideia "good" → retorna "Fail!"
// - Se houver 1 ou 2 ideias "good" → retorna "Publish!"
// - Se houver mais de 2 ideias "good" → retorna "I smell a series!"
func Well(ideias []string) string {

	var numBoasIdeias int // Cria uma variável que vai contar quantas ideias "good" existem

	// Percorre toda a lista de ideias usando "for range"
	// "_" ignora o índice e "ideia" representa o valor em cada posição
	for _, ideia := range ideias {
		// Se o valor atual for igual a "good"
		if ideia == "good" {
			numBoasIdeias++ // Incrementa o contador (somamos +1)
		}
	}

	// Agora, com base na contagem de boas ideias, tomamos a decisão:

	// Se não tiver nenhuma "good"
	if numBoasIdeias == 0 {
		return "Fail!" // Retorna "Fail!"
	}

	// Se tiver 1 ou 2 "good"
	if numBoasIdeias <= 2 {
		return "Publish!" // Retorna "Publish!"
	}

	// Se tiver mais de 2 "good"
	if numBoasIdeias > 2 {
		return "I smell a series!" // Retorna "I smell a series!"
	}

	// Esse return final é apenas uma segurança — o fluxo nunca deve chegar aqui.
	return ""
}

func main() {
	// 🧪 Testes práticos para entender o comportamento da função

	fmt.Println("Teste 1:", Well([]string{"bad", "bad", "bad"}))                // Esperado: Fail!
	fmt.Println("Teste 2:", Well([]string{"good", "bad", "bad"}))              // Esperado: Publish!
	fmt.Println("Teste 3:", Well([]string{"good", "bad", "good"}))             // Esperado: Publish!
	fmt.Println("Teste 4:", Well([]string{"good", "good", "good", "bad"}))     // Esperado: I smell a series!
	fmt.Println("Teste 5:", Well([]string{"good", "good", "good", "good"}))    // Esperado: I smell a series!
	fmt.Println("Teste 6:", Well([]string{}))                                  // Esperado: Fail! (lista vazia)
}
