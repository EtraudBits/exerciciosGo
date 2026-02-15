package main

import (
	"bufio"   // usado na linha 19 para ler a entrada do usuario de forma mais robusta.
	"fmt"     // usado na linha 56-57-59-61 para exibir mensagens e resultados formatados.
	"os"      // usado na linha 19 para acessar a entrada padrao do sistema.
	"strconv" // usado na linha 24-40 para converter a entrada de texto em numeros decimais (float64).
	"strings" // usado na linha 21-37 para limpar a entrada de texto, removendo espacos e quebras de linha.
)

func calcularGorgeta(totalContaCentavos int, porcentagemGorgeta int) int { // Funcao que calcula o valor da gorgeta em centavos.
	return totalContaCentavos * porcentagemGorgeta / 100
}

func formaPagamento(forma string, totalContaCentavos int) int { // funcao que calcula o valor do desconto em centavos.
	if forma == "dinheiro" || forma == "pix" { // Verifica se a forma de pagamento e dinheiro ou Pix.
		return totalContaCentavos * 10 / 100
	} 
	return 0 // Retorna 0 para outras formas de pagamento, sem desconto.
	}

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

func main() {

	reader := bufio.NewReader(os.Stdin) // Cria um leitor para a entrada padrao

	var err error // Declara a variavel de erro no escopo do main.
	
	totalContaCentavos := 0 // Inicializa a conta em centavos.
	for { // Loop para garantir uma conta valida.
		fmt.Print("Qual o valor da conta? ") // Exibe o prompt da conta.
		totalContaTexto, _ := reader.ReadString('\n') // Le a conta como texto.
		totalContaTexto = strings.TrimSpace(totalContaTexto) // Remove espacos e o \n do fim.

		totalContaCentavos = reaisParaCentavos(totalContaTexto) // Converte texto para centavos.
		if totalContaCentavos <= 0 { // Verifica erro na conversao.
			fmt.Println("Valor da conta inválido. Por favor insira um número!") // Mensagem de erro.
			continue // Volta para tentar novamente.
		}
		break // Sai do loop quando a conta e valida.
	}

	metodoPagamento := "" // Inicializa a forma de pagamento como string vazia.
	for { // Loop para garantir uma forma de pagamento valida.
		fmt.Print("Forma de pagamento? (1=dinheiro, 2=pix, 3=outros): ") // Exibe o prompt da forma de pagamento.
		metodoPagamento, err = reader.ReadString('\n') // Le a forma de pagamento como texto.
		if err != nil { // Verifica erro de leitura.
			fmt.Println("Erro ao ler a forma de pagamento. Tente novamente!") // Mensagem de erro.
			continue // Volta para tentar novamente.
		}
		metodoPagamento = strings.TrimSpace(metodoPagamento) // Remove espacos e o \n do fim.
		switch metodoPagamento { // Verifica a forma de pagamento usando switch para validar as opcoes.
		case "1":
			metodoPagamento = "dinheiro"
		case "2":
			metodoPagamento = "pix"
		case "3":
			metodoPagamento = "outros"
		default:
			fmt.Println("Forma de pagamento inválida. Use 1 (dinheiro), 2 (pix) ou 3 (outros).") // Mensagem de erro.
			continue // Volta para tentar novamente.
		}
		break // Sai do loop quando a forma de pagamento e valida.
	}

	porcentagemGorgeta := 10 // Gorjeta padrao para outros.
	if metodoPagamento == "dinheiro" || metodoPagamento == "pix" {
		porcentagemGorgeta = 15
	}

	gorgetaCentavos := calcularGorgeta(totalContaCentavos, porcentagemGorgeta)
	descontoCentavos := formaPagamento(metodoPagamento, totalContaCentavos)
	totalGeralCentavos := totalContaCentavos + gorgetaCentavos - descontoCentavos

	fmt.Printf("\n")
	fmt.Printf("Valor da gorgeta: R$ %.2f\n", float64(gorgetaCentavos)/100)
	fmt.Printf("\n")
	fmt.Printf("Desconto aplicado: R$ %.2f\n", float64(descontoCentavos)/100)
	fmt.Printf("\n")
	fmt.Printf("Valor total a pagar (conta + gorgeta - desconto): R$ %.2f\n", float64(totalGeralCentavos)/100)
	fmt.Printf("\n")

}