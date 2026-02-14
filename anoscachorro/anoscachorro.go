package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isNomeValido(nome string) bool {
	if nome == "" {
		return false
	}
	for _, r := range nome {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	//var nomeCao string
	//var idade int
	var err error

	var nomeCao string
	var entrada string
	for {
		fmt.Print("Qual o nome do cao?")
		//fmt.Scan(&nomeCao)
		entrada, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler o nome. Tente novamente!")
			continue
		}
		nomeCao = strings.TrimSpace(entrada)
		if !isNomeValido(nomeCao) {
			fmt.Println("Nome invalido. Use apenas letras e espacos.")
			continue
		}
		break
	}
	
	idadeHumana := 0
	for {
		fmt.Print("Qual a idade dele (Anos Humano)?")
		//fmt.Scan(&idade)
		idadeTexto, _ := reader.ReadString('\n')
		idadeTexto = strings.TrimSpace(idadeTexto)

		idadeHumana, err = strconv.Atoi(idadeTexto)
		if err != nil {
			fmt.Println("Idade inválida. Por favor insira um número inteiro!")
			continue
		}
		break
	}

	anosCao := idadeHumana * 7

	fmt.Printf("A idade do cao %s, é de %d anos de cao\n", nomeCao, anosCao)
}