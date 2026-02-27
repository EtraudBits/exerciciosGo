package main

import (
	"errors"
	"fmt"
)

// buscarConteudo simula uma requisição web que pode falhar
// Ela retorna a string do conteúdo E um possível erro
func buscarConteudo(url string) (string, error) {
	if url == "" {
		return "", errors.New("URL inválida")
	}
	// simulação de sucesso
	conteudo := "<html><body> conteúdo de " + url + " </body></html>"
	return conteudo, nil // sucesso: erro é nil
}

func main () {
	// capturamos os dois retornos usando a declração curta :=
	texto, err := buscarConteudo("http://exemplo.com")
	// ENGENHARIA GO: tratamos o erro imediatamente
	// Se erro não for 'nil', algo deu errado na RAM ou na rede
	if err != nil {
		fmt.Printf("Falha na arquitetura: %v\n", err)
		return // Encerra a função main com segurança
	}
	// Só chegamos aqui se err for nil, ou seja, o processo deu certo.
	fmt.Println("processamento:", texto)
}