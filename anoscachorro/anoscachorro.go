package main // Define o pacote principal do programa.

import ( // Inicia o bloco de imports.
	"bufio"   // Leitura com buffer do teclado.
	"fmt"     // Funcoes de entrada/saida formatadas.
	"os"      // Acesso a recursos do sistema operacional.
	"strconv" // Conversao de string para numero.
	"strings" // Funcoes auxiliares para string.
	"unicode" // Classificacao de caracteres.
) // Finaliza o bloco de imports.

func isNomeValido(nome string) bool { // Funcao que valida o nome do cao.
	if nome == "" { // Verifica se a string esta vazia.
		return false // Retorna falso se estiver vazia.
	}
	for _, r := range nome { // Percorre cada caractere do nome.
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) { // Permite apenas letras e espacos.
			return false // Retorna falso se encontrar algo invalido.
		}
	}
	return true // Retorna verdadeiro quando tudo esta valido.
}

func calculaAnosCao(idadeHumana float64) float64 { // Calcula a idade canina pela regra 15/9/5.
	if idadeHumana <= 0 { // Trata valores zero ou negativos.
		return 0 // Retorna 0 para idade invalida.
	}
	if idadeHumana <= 1 { // Primeiro ano vale 15.
		return idadeHumana * 15 // Proporcional para decimais no primeiro ano.
	}
	if idadeHumana <= 2 { // Segundo ano vale 9.
		return 15 + (idadeHumana-1)*9 // Soma do primeiro ano mais o restante.
	}
	return 24 + (idadeHumana-2)*5 // 15 + 9 e depois 5 por ano.
}

func main() { // Ponto de entrada do programa.

	reader := bufio.NewReader(os.Stdin) // Cria um leitor para a entrada padrao.

	//var nomeCao string // Declaracao antiga (comentada).
	//var idade int // Declaracao antiga (comentada).
	
	var err error // Declara a variavel de erro no escopo do main.

	var nomeCao string // Declara a variavel do nome.
	var entrada string // Declara a variavel para entrada de texto.
	
	for { // Loop para garantir um nome valido.
		fmt.Print("Qual o nome do cao?") // Exibe o prompt do nome.
		//fmt.Scan(&nomeCao) // Leitura simples (comentada).
		entrada, err = reader.ReadString('\n') // Le ate a quebra de linha.
		if err != nil { // Verifica erro de leitura.
			fmt.Println("Erro ao ler o nome. Tente novamente!") // Mensagem de erro.
			continue // Volta para tentar novamente.
		}
		nomeCao = strings.TrimSpace(entrada) // Remove espacos e o \n do fim.
		if !isNomeValido(nomeCao) { // Valida se o nome e aceito.
			fmt.Println("Nome invalido. Use apenas letras e espacos.") // Mensagem de erro.
			continue // Volta para tentar novamente.
		}
		break // Sai do loop quando o nome e valido.
	}
	
	idadeHumana := 0.0 // Inicializa a idade humana como decimal.
	for { // Loop para garantir uma idade valida.
		fmt.Print("Qual a idade dele (Anos Humano)?") // Exibe o prompt da idade.
		//fmt.Scan(&idade) // Leitura simples (comentada).
		idadeTexto, _ := reader.ReadString('\n') // Le a idade como texto.
		idadeTexto = strings.TrimSpace(idadeTexto) // Remove espacos e o \n do fim.

		idadeHumana, err = strconv.ParseFloat(idadeTexto, 64) // Converte texto para decimal.
		if err != nil { // Verifica erro na conversao.
			fmt.Println("Idade inválida. Por favor insira um número!") // Mensagem de erro.
			continue // Volta para tentar novamente.
		}
		break // Sai do loop quando a idade e valida.
	}

	anosCao := calculaAnosCao(idadeHumana) // Calcula a idade do cao em anos caninos.

	fmt.Printf("A idade do cao %s, é de %.2f anos de cao\n", nomeCao, anosCao) // Exibe o resultado.
	fmt.Println("até 1 ano: idade × 15 anos de cao")
	fmt.Println("de 1 a 2 anos: 15 + (idade - 1) × 9 anos de cao")
	fmt.Println("acima de 2 anos: 24 + (idade - 2) × 5 anos de cao")
} // Finaliza a funcao main.