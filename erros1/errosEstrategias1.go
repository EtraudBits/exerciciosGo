// 5 extratergias de tratamento de erros.
// 1 Programar
// 2 Repetir
// 3 Encerrar
// 4 logar e continuar
// ignorar (Raro)

package main

import (
	"fmt"
	"net/http" // pacote para simular resisições reais.
)

// buscarHTML ilustra a Estratégia 1: programação com contexto de erro.
func buscarHTML(url string) (*http.Response, error) { // *http.Response é o tipo de retorno esperado (ponteiro para uma resposta HTTP), e error é o tipo de erro que pode ocorrer.
	resp, err := http.Get(url)  // Simula uma requisição HTTP para a url fornecida pelo usuario. O resultado é armazenado em resp, e qualquer erro é armazenado em err.

	if err != nil { // se não for nulo (nil), significa que ocorreu um erro durante a requisição e a mensagem de erro dentro do bloco if será executado...
		// em vez de apenas retornar 'err', adicionamos contexto (arquitetura de erro) para tornar a mensagem mais informativa.
		return nil, fmt.Errorf("falha na rede ao acessar %s: %v\n", url, err)
	}
	return resp, nil // Se não houver erro, retorna a resposta e nil para indicar que não houve erro.
}

func main() {
	url := "http://site-inexistente.com" // Variável que armazena a URL que será acessada. Neste caso, é uma URL ficticia que provavelmente causará um erro de rede.
	_, err := buscarHTML(url) // Chama a função buscarHTML com a URL fornecida e armazena o resultado em _ (ignorado) e o erro em err.
	if err != nil { // novamente, verifica se ocorre um erro ao chamar buscarHTML. Se err não for nulo (nil), significa que houve um erro e a mensagem de erro será mostrada.
		// Estretegia 4: logar o erro e parar esta execusão especifica, mas continuar a execução do programa.
		fmt.Printf("LOG DE ENGENHARIA: %v\n", err) // Imprime a mensagem de erro formatada no console, incluindo o contexto adicional fornecido pela a função buscarHTML.
		return // encerra a execução do main, mas o programa em si pode continuar rodando se houver outra parte do código após este bloco.
	}
	fmt.Println("Sucesso ao carregar a página!") // se a requisição for bem-sucedida, esta mensagem será exibida no console.
}