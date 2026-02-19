package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Struct para representar um aluno, com nome e nota.
type aluno struct { 
	nome string
	notas [4]float64
}
// Struct para representar o registro de notas, que contém um slice de alunos.
type registroNotas struct {
	alunos []aluno
}

func (r *registroNotas) adicionarAluno(nome string, notas [4]float64) { // Metodo para adicionar um aluno ao registro de notas.
	r.alunos = append(r.alunos, aluno{nome: nome, notas: notas}) // Adiciona um novo aluno ao slice de alunos do registro.
}

func (r *registroNotas) buscarAluno(nome string) *aluno { // Metodo para buscar um aluno pelo nome no registro de notas.
	for i, a := range r.alunos {
		if a.nome == nome {
			return &r.alunos[i]
		}
	}
	return nil // Retorna nil se o aluno não for encontrado.
}

func (r *registroNotas) atualizarNotas(nome string, novasNotas [4]float64) bool { // Metodo para atualizar as notas de um aluno no registro de notas.
	for i, a := range r.alunos { // Loop para encontrar o aluno a ser atualizado.
		if a.nome == nome { // Verifica se o nome do aluno atual é o que queremos atualizar.
			r.alunos[i].notas = novasNotas // Atualiza as notas do aluno encontrado.
			return true // Retorna true para indicar que a atualização foi bem-sucedida.
		}
	}
	return false // Retorna false se o aluno não for encontrado.
}

func (r *registroNotas) incluirNotaAluno(nome string, numeroNota int, nota float64) bool {
	if numeroNota < 1 || numeroNota > 4 {
		return false
	}

	for i, a := range r.alunos {
		if a.nome == nome {
			r.alunos[i].notas[numeroNota-1] = nota
			return true
		}
	}
	return false
}

func (r *registroNotas) calcularMediaAluno(nome string) (float64, bool) {
	for _, a := range r.alunos {
		if a.nome == nome {
			soma := 0.0
			for _, nota := range a.notas {
				soma += nota
			}
			return soma / 4.0, true
		}
	}
	return 0.0, false
}

func (r *registroNotas) removerAluno(nome string) bool { // Metodo para remover um aluno do registro de notas.
	for i, a := range r.alunos { // Loop para encontrar o aluno a ser removido.
		if a.nome == nome { // Verifica se o nome do aluno atual é o que queremos remover.
			r.alunos = append(r.alunos[:i], r.alunos[i+1:]...) // Remove o aluno do slice de alunos usando slicing para criar um novo slice sem o aluno removido.
			return true // Retorna true para indicar que a remoção foi bem-sucedida.
		}}
	return false // Retorna false se o aluno não for encontrado.
}

func (r *registroNotas) exibirAlunos() { // Metodo para exibir os alunos e suas notas no registro de notas.
	fmt.Println("Alunos e Notas:") // Exibe o titulo da lista de alunos e notas.
	for _, a := range r.alunos { // Loop para exibir cada aluno e sua nota.
		soma := 0.0
		for _, nota := range a.notas {
			soma += nota
		}
		media := soma / 4.0
		fmt.Printf("Nome: %s, Notas: [%.2f, %.2f, %.2f, %.2f], Média: %.2f\n", a.nome, a.notas[0], a.notas[1], a.notas[2], a.notas[3], media)
	}
}

func lerNota(reader *bufio.Reader, mensagem string) (float64, bool) {
	fmt.Print(mensagem)
	notaTexto, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a nota.")
		return 0, false
	}

	notaTexto = strings.TrimSpace(notaTexto)
	nota, err := strconv.ParseFloat(notaTexto, 64)
	if err != nil {
		fmt.Println("Nota inválida. A nota deve ser um número.")
		return 0, false
	}

	return nota, true
}

func nomeValido(nome string) bool {
	nome = strings.TrimSpace(nome)
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

func confirmarAcao(reader *bufio.Reader, pergunta string) bool {
	fmt.Print(pergunta)
	resposta, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler confirmação. Operação cancelada.")
		return false
	}

	resposta = strings.TrimSpace(strings.ToLower(resposta))
	return resposta == "s" || resposta == "sim"
}



func main() {
	reader := bufio.NewReader(os.Stdin) // Cria um leitor para a entrada padrao

	var registro registroNotas // Cria uma variavel do tipo registroNotas para armazenar os alunos e suas notas.

	
	for { // Loop para garantir uma opcao valida.
		fmt.Println("\nEscolha uma opcao:") // Exibe as opcoes para o usuario.
		fmt.Println("1. Adicionar aluno") // Opcao para adicionar um aluno.
		fmt.Println("2. Buscar aluno") // Opcao para buscar um aluno.
		fmt.Println("3. Atualizar as 4 notas") // Opcao para atualizar as 4 notas de um aluno.
		fmt.Println("4. Calcular media do aluno") // Opcao para calcular a media de um aluno.
		fmt.Println("5. Remover aluno") // Opcao para remover um aluno do registro.
		fmt.Println("6. Exibir alunos") // Opcao para exibir os alunos e suas notas.
		fmt.Println("7. Incluir nota (1 a 4) em aluno existente") // Opcao para incluir nota em aluno existente.
		fmt.Println("8. Sair") // Opcao para sair do programa.

		opcao, err := reader.ReadString('\n') // Le a opcao escolhida pelo usuario.
		if err != nil { // Verifica erro de leitura.
			fmt.Println("Erro ao ler a opcao. Tente novamente!") // Mensagem de erro.
			continue // Volta para tentar novamente.
		}
		opcao = strings.TrimSpace(opcao) // Remove espacos e o \n do fim da opcao.

		switch opcao { // Verifica a opcao escolhida usando switch para validar as opcoes.
		case "1":
			fmt.Print("Digite o nome do aluno: ") // Prompt para o nome do aluno.
			nome, _ := reader.ReadString('\n') // Le o nome do aluno como texto.
			nome = strings.TrimSpace(nome) // Remove espacos e o \n do fim do nome.
			if !nomeValido(nome) {
				fmt.Println("Aluno inválido. Digite um nome com texto (apenas letras e espaços).")
				continue
			}

			var notas [4]float64
			valido := true
			for i := 0; i < 4; i++ {
				nota, ok := lerNota(reader, fmt.Sprintf("Digite a nota %d do aluno: ", i+1))
				if !ok {
					valido = false
					break
				}
				notas[i] = nota
			}

			if !valido {
				continue
			}

			registro.adicionarAluno(nome, notas) // Adiciona o aluno ao registro de notas.
			fmt.Println("Aluno adicionado com sucesso!") // Mensagem de sucesso.

		case "2":
			fmt.Print("Digite o nome do aluno a ser buscado: ") // Prompt para o nome do aluno a ser buscado.
			nome, _ := reader.ReadString('\n') // Le o nome do aluno como texto.
			nome = strings.TrimSpace(nome) // Remove espacos e o \n do fim do nome.

			alunoEncontrado := registro.buscarAluno(nome) // Busca o aluno no registro de notas.
			if alunoEncontrado != nil { // Verifica se o aluno foi encontrado.
				media, _ := registro.calcularMediaAluno(nome)
				fmt.Printf("Aluno encontrado: Nome: %s, Notas: [%.2f, %.2f, %.2f, %.2f], Média: %.2f\n", alunoEncontrado.nome, alunoEncontrado.notas[0], alunoEncontrado.notas[1], alunoEncontrado.notas[2], alunoEncontrado.notas[3], media)
			} else {
				fmt.Println("Aluno não encontrado.") // Mensagem de erro se o aluno não for encontrado.
			}

		case "3":
			if !confirmarAcao(reader, "Deseja alterar a nota de um aluno? (s/n): ") {
				fmt.Println("Alteração cancelada.")
				continue
			}

			fmt.Print("Digite o nome do aluno para atualizar a nota: ") // Prompt para o nome do aluno a ser atualizado.
			nome, _ := reader.ReadString('\n') // Le o nome do aluno como texto.
			nome = strings.TrimSpace(nome) // Remove espacos e o \n do fim do nome.
			if !nomeValido(nome) {
				fmt.Println("Aluno inválido. Digite um nome com texto (apenas letras e espaços).")
				continue
			}

			var notas [4]float64
			valido := true
			for i := 0; i < 4; i++ {
				nota, ok := lerNota(reader, fmt.Sprintf("Digite a nova nota %d do aluno: ", i+1))
				if !ok {
					valido = false
					break
				}
				notas[i] = nota
			}

			if !valido {
				continue
			}

			if registro.atualizarNotas(nome, notas) { // Tenta atualizar as notas do aluno no registro de notas.
				fmt.Println("Nota atualizada com sucesso!") // Mensagem de sucesso se a atualização for bem-sucedida.
			} else {
				fmt.Println("Aluno não encontrado. Não foi possível atualizar a nota.") // Mensagem de erro se o aluno não for encontrado para atualização.
			}

		case "4":
			fmt.Print("Digite o nome do aluno para calcular a média: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)
			if !nomeValido(nome) {
				fmt.Println("Aluno inválido. Digite um nome com texto (apenas letras e espaços).")
				continue
			}

			media, ok := registro.calcularMediaAluno(nome)
			if !ok {
				fmt.Println("Aluno não encontrado.")
				continue
			}

			fmt.Printf("A média do aluno %s é: %.2f\n", nome, media)

		case "5":
			if !confirmarAcao(reader, "Deseja remover um aluno? (s/n): ") {
				fmt.Println("Remoção cancelada.")
				continue
			}

			fmt.Print("Digite o nome do aluno a ser removido: ") // Prompt para o nome do aluno a ser removido.
			nome, _ := reader.ReadString('\n') // Le o nome do aluno como texto.
			nome = strings.TrimSpace(nome) // Remove espacos e o \n do fim do nome.

			if registro.removerAluno(nome) { // Tenta remover o aluno do registro de notas.
				fmt.Println("Aluno removido com sucesso!") // Mensagem de sucesso se a remoção for bem-sucedida.
			} else {
				fmt.Println("Aluno não encontrado. Não foi possível remover.") // Mensagem de erro se o aluno não for encontrado para remoção.
			}

		case "6":
			registro.exibirAlunos() // Exibe os alunos e suas notas no registro de notas.

		case "7":
			fmt.Print("Digite o nome do aluno para incluir nova nota: ")
			nome, _ := reader.ReadString('\n')
			nome = strings.TrimSpace(nome)
			if !nomeValido(nome) {
				fmt.Println("Aluno inválido. Digite um nome com texto (apenas letras e espaços).")
				continue
			}

			fmt.Print("Escolha qual nota deseja incluir (1, 2, 3 ou 4): ")
			numeroTexto, _ := reader.ReadString('\n')
			numeroTexto = strings.TrimSpace(numeroTexto)

			numeroNota, err := strconv.Atoi(numeroTexto)
			if err != nil || numeroNota < 1 || numeroNota > 4 {
				fmt.Println("Opção inválida. Escolha um número entre 1 e 4.")
				continue
			}

			novaNota, ok := lerNota(reader, "Digite o valor da nota: ")
			if !ok {
				continue
			}

			if registro.incluirNotaAluno(nome, numeroNota, novaNota) {
				fmt.Println("Nota incluída com sucesso!")
			} else {
				fmt.Println("Aluno não encontrado. Não foi possível incluir a nota.")
			}

		case "8":
			if !confirmarAcao(reader, "Deseja sair do programa? (s/n): ") {
				fmt.Println("Saída cancelada.")
				continue
			}

			fmt.Println("Saindo do programa...") // Mensagem de saida.
			os.Exit(0) // Encerra o programa.

		default: // Caso para opção inválida.
			fmt.Println("Opção inválida. Por favor escolha uma opção entre 1 e 8.") // Mensagem de erro para opção inválida.
		}
	}
	
}