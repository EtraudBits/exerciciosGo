// Erros Estruturados - Parte 2
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Estratefgia 2: retentativa (Retry)
// tenta baixa a página 3 vezes antes de desistir
func buscarComRetry(url string) (*http.Response, error) {
	const timeout = 1 * time.Minute // define um prazo para as tentativas, neste caso, 1 minuto
	deadline := time.Now(). Add(timeout) // calcula o prazo final para as tentativas, neste caso, 1 minuto a partir do momento atual

	for tentativas := 0; time.Now().Before(deadline); tentativas++ { // for que continua enquanto o tempo atual for antes do prazo final
		resp, err := http.Get(url) // tenta baixar a página (url) usando http.Get na função buscarComRetry
		if err == nil { // se a resposta for bem-sucedida (sem erros), retorna a resposta e nil (indicando sucesso)
			return resp, nil // sucesso
		}
		log.Printf("Servidor não respondeu (%s), tentando novamente...", err) // se houver um erro, loga a mensagem de erro e continua para a próxima tentativa
		time.Sleep(time.Second << (tentativas)) // espera exponencial: 1s, 2s, 4s, etc.
	}
	// Estrategia 1: Propagação
	return nil, fmt.Errorf("servidor %s não responde após 1 minuto", url) // se o prazo final for atingido sem sucesso, retorna um erro formatado indicando que o servidor não respondeu após 1 minuto
}

func main() {
	_, err := buscarComRetry("http://site-instavel.com") // chama a função buscarComRetry com a URL "http://site-instavel.com" e armazena o erro retornado (se houver) na variável err, "_" é usado para ignorar a resposta bem-sucedida, pois o foco aqui é lidar com o erro	
	if err != nil { // se houver um erro, ou seja, se err for diferente de nil, entra neste bloco
		// Estrategia 3: interrupção (log.Fatal encerra o programa)
		log.Fatalf("Erro crítico no crowler: %v", err) // loga a mensagem de erro crítica e encerra o programa usando log.Fatalf, que formata a mensagem de erro e termina a execução do programa
	}
}