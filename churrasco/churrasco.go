package main // define o pacote principal do programa
// linha em branco comentada para manter separacao visual
import ( // inicia o bloco de imports
	"fmt"     // funcoes de entrada e saida
	"strconv" // conversao de string para inteiro
	"strings" // funcoes utilitarias de string
) // encerra o bloco de imports
// linha em branco comentada para manter separacao visual
func main() { // funcao principal do programa
	var pessoa int // armazena a quantidade de pessoas
	var precoCarneCentavos int // armazena o preco por kg em centavos
	var precoCarneReais string // armazena o preco por kg em reais como texto

	fmt.Print("Quantas pessoas vao ao churrasco?") // exibe a pergunta
	fmt.Scan(&pessoa) // le a quantidade de pessoas

	fmt.Print("Qual o preco da carne (kg) em reais? (ex: 12.50)") // exibe a pergunta do preco
	fmt.Scan(&precoCarneReais) // le o preco em reais como texto

	precoCarneCentavos = reaisParaCentavos(precoCarneReais) // converte o preco para centavos

	const mediaPorPessoaGramas = 500 // define 0.5 kg por pessoa em gramas

	quantidadeCarneGramas := pessoa * mediaPorPessoaGramas // calcula a quantidade total em gramas
	custoTotalCentavos := quantidadeCarneGramas * precoCarneCentavos / 1000 // calcula o custo total em centavos
	quantidadeCarneKg := float64(quantidadeCarneGramas) / 1000.0 // converte gramas para kg
	
	fmt.Printf("a quantidade de carne para %d pessoas no churrasco é: %.2f kg\n", pessoa, quantidadeCarneKg) // imprime a quantidade de carne
	fmt.Printf("O custo total da carne para %d pessoas no churrasco é: R$ %d.%02d\n", pessoa, custoTotalCentavos/100, custoTotalCentavos%100) // imprime o custo total
} // encerra a funcao main




func reaisParaCentavos(valor string) int { // converte string em reais para centavos
	valor = strings.TrimSpace(valor) // remove espacos no inicio e fim
	partes := strings.SplitN(valor, ".", 2) // separa parte inteira e decimal
	reais, _ := strconv.Atoi(partes[0]) // converte a parte inteira em inteiro
	centavos := 0 // inicia os centavos com zero
	if len(partes) == 2 { // verifica se existe parte decimal
		frac := partes[1] // guarda a parte decimal
		if len(frac) == 1 { // se so houver um digito
			frac += "0" // completa com zero
		} else if len(frac) > 2 { // se tiver mais de dois digitos
			frac = frac[:2] // usa apenas os dois primeiros
		} // encerra o ajuste da parte decimal
		centavos, _ = strconv.Atoi(frac) // converte a parte decimal em inteiro
	} // encerra o bloco do decimal
	return reais*100 + centavos // retorna o total em centavos
} // encerra a funcao reaisParaCentavos