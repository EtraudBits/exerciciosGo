// aqui Go lida com a leitura de caracteres individuais ou blocos de uma pagina até o fim

package main

import (
	"fmt"
	"io" // pacote io é usado para lidar com operações de entrada/saída, como leitura e escrita de dados. Ele fornece a interface io.Reader, que é implementada por tipos que podem ser lidos, como strings.NewReader. Neste código, io.EOF é usado para detectar quando a leitura atingiu o final do arquivo (ou seja, quando não há mais dados para ler).
	"strings"
)

func main() {
	// simulando o corpo de uma página HTML.
	leitor := strings.NewReader("<html>Link1, Link2</html>")

	//criamos um buffer (pequeno espaço na RAM) para ler 1 byte por vez
	p := make([]byte, 1) 

	for {
		_, err := leitor.Read(p) // tenta ler o próximo byte
		// engenharia: verificamos se chegamos ao fim (Estrategia EOF) primeiro, antes de verificar outros erros.
		if err == io.EOF {
			fmt.Println("\n -- Fim da Leitura: Página processada com sucesso --")
			break // sai do loop de leitura de forma limpa, indicando que a página foi processada com sucesso
		}
		// verificamos se houve um erro REAL (ex.: corrupção de dados).
		if err != nil {
			fmt.Printf("Erro catastrofico na leitura: %v\n", err)
			return
		}
		// se err foi nil, processamos o byte lido (neste caso, apenas imprimimos o caractere)
		fmt.Printf("%c", p[0])
	}
}