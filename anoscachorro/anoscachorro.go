package main // Define o pacote principal do programa.

import ( // Inicia o bloco de imports.
	"bufio"   // Leitura com buffer do teclado.
	"fmt"     // Funcoes de entrada/saida formatadas.
	"os"      // Acesso a recursos do sistema operacional.
	"strconv" // Conversao de string para numero.
	"strings" // Funcoes auxiliares para string.
) // Finaliza o bloco de imports.

func main() { // Ponto de entrada do programa.

	reader := bufio.NewReader(os.Stdin) // Cria um leitor para a entrada padrão.

	//var nomeCao string // Exemplo antigo de declaracao (comentado).
	//var idade int // Exemplo antigo de declaracao (comentado).

	fmt.Print("Qual o nome do cao?") // Mostra o prompt do nome.
	//fmt.Scan(&nomeCao) // Leitura simples (comentada).
	nomeCao, err := reader.ReadString('\n') // Le o nome ate a quebra de linha.
	nomeCao = strings.TrimSpace(nomeCao) // Remove espacos e o \n do fim.
	
	idadeHumana := 0 // Declara fora do loop para usar no calculo final.
	for { // Loop para garantir uma idade valida.
		fmt.Print("Qual a idade dele (Anos Humano)?") // Mostra o prompt da idade.
		//fmt.Scan(&idade) // Leitura simples (comentada).
		idadeTexto, _ := reader.ReadString('\n') // Le a idade como texto.
		idadeTexto = strings.TrimSpace(idadeTexto) // Remove espacos e o \n do fim.

		idadeHumana, err = strconv.Atoi(idadeTexto) // Converte idade para inteiro.
		if err != nil { // Verifica se houve erro na conversao.
			fmt.Println("Idade inválida. Por favor insira um número inteiro!") // Mensagem de erro.
			continue // Volta para o início do loop para tentar novamente.
		}
		break // Sai do loop se a conversao foi bem-sucedida.
	}

	anosCao := idadeHumana * 7 // Calcula idade do cao em anos caninos.

	

	fmt.Printf("A idade do cao %s, é de %d anos de cao\n", nomeCao, anosCao) // Exibe o resultado.
}