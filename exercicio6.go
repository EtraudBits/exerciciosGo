package main // Pacote principal: permite executar o programa com `go run`

// Importa pacotes necessários:
// - fmt para imprimir no terminal
// - strings para usar strings.Repeat (função pronta que repete strings)
import (
	"fmt"
	"strings"
)

// RepeatStrLoop repete a string 'value' 'repetitions' vezes usando um loop for.
// Esta versão é didática: mostra a lógica passo a passo.
func RepeatStrLoop(repetitions int, value string) string {
	// Declara a variável que guardará o resultado acumulado.
	// Deve ficar fora do laço para manter o valor entre iterações.
	var repeatString string

	// Inicia o laço: contador i começa em 0; roda enquanto i < repetitions; incrementa i a cada volta.
	for i := 0; i < repetitions; i++ {
		// Concatena a string atual (value) ao acumulador.
		// Exemplo: "" + "Go" => "Go"; depois "Go" + "Go" => "GoGo", etc.
		repeatString += value

		// Nota: poderia usar `repeatString = repeatString + value`, que é equivalente.
	}

	// Depois que o loop termina, retornamos o string acumulado.
	return repeatString
}

// RepeatStrBuilt repete a string usando a função pronta do pacote strings.
// Esta é a forma recomendada em código real — mais curta e geralmente mais eficiente.
func RepeatStrBuilt(repetitions int, value string) string {
	// strings.Repeat recebe (valor string, n int) e retorna o resultado.
	// Se repetitions for <= 0, strings.Repeat retorna "" automaticamente.
	return strings.Repeat(value, repetitions)
}

func main() {
	// Exemplos para testar as duas funções e ver se ambas produzem o mesmo resultado.

	// Exemplo 1
	reps1 := 3               // número de repetições desejadas
	val1 := "Go"             // string a ser repetida
	// Chama a versão com loop e imprime o resultado esperado
	fmt.Println("Loop ->", RepeatStrLoop(reps1, val1))   // Esperado: "GoGoGo"
	// Chama a versão com função pronta e imprime o resultado esperado
	fmt.Println("Built ->", RepeatStrBuilt(reps1, val1)) // Esperado: "GoGoGo"

	// Exemplo 2: repetições zero (caso limite)
	reps2 := 0
	val2 := "Hi"
	// Ambas as funções devem retornar string vazia quando repetitions == 0
	fmt.Println("Loop (0) ->", RepeatStrLoop(reps2, val2))   // Esperado: ""
	fmt.Println("Built (0) ->", RepeatStrBuilt(reps2, val2)) // Esperado: ""

	// Exemplo 3: string vazia (outra borda)
	reps3 := 5
	val3 := ""
	// Repetir string vazia deve sempre resultar em string vazia
	fmt.Println("Loop (empty) ->", RepeatStrLoop(reps3, val3))   // Esperado: ""
	fmt.Println("Built (empty) ->", RepeatStrBuilt(reps3, val3)) // Esperado: ""

	// Exemplo 4: caractere especial e repetições maiores
	reps4 := 4
	val4 := "🙂"
	// Testando com caracteres unicode (em Go isso funciona corretamente com strings)
	fmt.Println("Loop (emoji) ->", RepeatStrLoop(reps4, val4))   // Esperado: "🙂🙂🙂🙂"
	fmt.Println("Built (emoji) ->", RepeatStrBuilt(reps4, val4)) // Esperado: "🙂🙂🙂🙂"
}
