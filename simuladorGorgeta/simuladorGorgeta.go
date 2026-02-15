package main

import (
	"bufio"   // usado na linha 19 para ler a entrada do usuario de forma mais robusta.
	"fmt"     // usado na linha 56-57-59-61 para exibir mensagens e resultados formatados.
	"os"      // usado na linha 19 para acessar a entrada padrao do sistema.
	"strconv" // usado na linha 24-40 para converter a entrada de texto em numeros decimais (float64).
	"strings" // usado na linha 21-37 para limpar a entrada de texto, removendo espacos e quebras de linha.
)

func calcularGorgeta(totalConta, porcentagemGorgeta float64) float64 { // Funcao que calcula o valor da gorgeta com base na conta e porcentagem.
	gorgeta := totalConta * (porcentagemGorgeta / 100) // Calcula a gorgeta multiplicando a conta pela porcentagem convertida para decimal.
	return gorgeta // Retorna o valor da gorgeta calculada.
}

func formaPagamento(forma string, totalConta float64) float64 { // funcao que calcula o valor do desconto com base na conta.
	if forma == "dinheiro" || forma == "pix" { // Verifica se a forma de pagamento e dinheiro ou Pix.
		return totalConta * 0.10 // Aplica um desconto de 10% para pagamento em dinheiro ou Pix.
	} 
	return 0 // Retorna 0 para outras formas de pagamento, sem desconto.
	}



func main() {

	reader := bufio.NewReader(os.Stdin) // Cria um leitor para a entrada padrao

	var err error // Declara a variavel de erro no escopo do main.
	
	totalConta := 0.0 // Inicializa a conta como decimal.
	for { // Loop para garantir uma conta valida.
		fmt.Print("Qual o valor da conta? ") // Exibe o prompt da conta.
		totalContaTexto, _ := reader.ReadString('\n') // Le a conta como texto.
		totalContaTexto = strings.TrimSpace(totalContaTexto) // Remove espacos e o \n do fim.

		totalConta, err = strconv.ParseFloat(totalContaTexto, 64) // Converte texto para decimal.
		if err != nil { // Verifica erro na conversao.
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

	porcentagemGorgeta := 10.0 // Gorjeta padrao para outros.
	if metodoPagamento == "dinheiro" || metodoPagamento == "pix" {
		porcentagemGorgeta = 15.0
	}

	gorgeta := calcularGorgeta(totalConta, porcentagemGorgeta)
	desconto := formaPagamento(metodoPagamento, totalConta)
	totalGeral := totalConta + gorgeta - desconto

	fmt.Printf("\n")
	fmt.Printf("Valor da gorgeta: R$ %.2f\n", gorgeta)
	fmt.Printf("\n")
	fmt.Printf("Desconto aplicado: R$ %.2f\n", desconto)
	fmt.Printf("\n")
	fmt.Printf("Valor total a pagar (conta + gorgeta - desconto): R$ %.2f\n", totalGeral)
	fmt.Printf("\n")

}