package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct para representar um aluno, com nome e nota.
type aluno struct { 
	nome string
	nota float64
}
// Struct para representar o registro de notas, que contém um slice de alunos.
type registroNotas struct {
	alunos []aluno
}

func (r *registroNotas) adicionarAluno(nome string, nota float64) { // Metodo para adicionar um aluno ao registro de notas.
	r.alunos = append(r.alunos, aluno{nome: nome, nota: nota}) // Adiciona um novo aluno ao slice de alunos do registro.
}

func (r *registroNotas) buscarAluno(nome string) *aluno { // Metodo para buscar um aluno pelo nome no registro de notas.
	alunos := map[string]aluno{} // Cria um mapa para armazenar os alunos por nome para busca eficiente.
	for _, a := range r.alunos { // Loop para preencher o mapa com os alunos do registro.
		alunos[a.nome] = a // Adiciona o aluno ao mapa usando o nome como chave.
	}
	if a, ok := alunos[nome]; ok { // Verifica se o aluno existe no mapa.
		return &a // Retorna o ponteiro para o aluno encontrado.
	}
	return nil // Retorna nil se o aluno não for encontrado.
}

func (r *registroNotas) atualizarNotas(nome string, novaNota float64) bool { // Metodo para atualizar a nota de um aluno no registro de notas.
	for i, a := range r.alunos { // Loop para encontrar o aluno a ser atualizado.
		if a.nome == nome { // Verifica se o nome do aluno atual é o que queremos atualizar.
			r.alunos[i].nota = novaNota // Atualiza a nota do aluno encontrado.
			return true // Retorna true para indicar que a atualização foi bem-sucedida.
		}
	}
	return false // Retorna false se o aluno não for encontrado.
}

func (r *registroNotas) calcularMedia() float64 { // Metodo para calcular a media das notas dos alunos no registro.
	if len(r.alunos) == 0 { // Verifica se não há alunos no registro para evitar divisão por zero.
		return 0.0 // Retorna 0.0 como media se não houver alunos.
	}
	var soma float64 // Variavel para armazenar a soma das notas dos alunos.
	for _, a := range r.alunos { // Loop para somar as notas dos alunos.
		soma += a.nota // Adiciona a nota do aluno à soma.
	}
	return soma / float64(len(r.alunos)) // Retorna a media das notas dos alunos.
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
		fmt.Printf("Nome: %s, Nota: %.2f\n", a.nome, a.nota) // Exibe o nome e a nota do aluno formatada com duas casas decimais.
	}
}



func main() {
	reader := bufio.NewReader(os.Stdin) // Cria um leitor para a entrada padrao

	var registro registroNotas // Cria uma variavel do tipo registroNotas para armazenar os alunos e suas notas.

	
	for { // Loop para garantir uma opcao valida.
		fmt.Println("\nEscolha uma opcao:") // Exibe as opcoes para o usuario.
		fmt.Println("1. Adicionar aluno") // Opcao para adicionar um aluno.
		fmt.Println("2. Buscar aluno") // Opcao para buscar um aluno.
		fmt.Println("3. Atualizar nota") // Opcao para atualizar a nota de um aluno.
		fmt.Println("4. Calcular media") // Opcao para calcular a media das notas dos alunos.
		fmt.Println("5. Remover aluno") // Opcao para remover um aluno do registro.
		fmt.Println("6. Exibir alunos") // Opcao para exibir os alunos e suas notas.
		fmt.Println("7. Sair") // Opcao para sair do programa.

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

			fmt.Print("Digite a nota do aluno: ") // Prompt para a nota do aluno.
			notaTexto, _ := reader.ReadString('\n') // Le a nota do aluno como texto.
			notaTexto = strings.TrimSpace(notaTexto) // Remove espacos e o \n do fim da nota.

			nota, err := strconv.ParseFloat(notaTexto, 64) // Converte a nota de texto para float64.
			if err != nil { // Verifica erro na conversao da nota.
				fmt.Println("Nota inválida. Por favor insira um número!") // Mensagem de erro.
				continue // Volta para tentar novamente.
			}

			registro.adicionarAluno(nome, nota) // Adiciona o aluno ao registro de notas.
			fmt.Println("Aluno adicionado com sucesso!") // Mensagem de sucesso.

		case "2":
			fmt.Print("Digite o nome do aluno a ser buscado: ") // Prompt para o nome do aluno a ser buscado.
			nome, _ := reader.ReadString('\n') // Le o nome do aluno como texto.
			nome = strings.TrimSpace(nome) // Remove espacos e o \n do fim do nome.

			alunoEncontrado := registro.buscarAluno(nome) // Busca o aluno no registro de notas.
			if alunoEncontrado != nil { // Verifica se o aluno foi encontrado.
				fmt.Printf("Aluno encontrado: Nome: %s, Nota: %.2f\n", alunoEncontrado.nome, alunoEncontrado.nota) // Exibe o nome e a nota do aluno encontrado.
			} else {
				fmt.Println("Aluno não encontrado.") // Mensagem de erro se o aluno não for encontrado.
			}

		case "3":
			fmt.Print("Digite o nome do aluno para atualizar a nota: ") // Prompt para o nome do aluno a ser atualizado.
			nome, _ := reader.ReadString('\n') // Le o nome do aluno como texto.
			nome = strings.TrimSpace(nome) // Remove espacos e o \n do fim do nome.

			fmt.Print("Digite a nova nota do aluno: ") // Prompt para a nova nota do aluno.
			notaTexto, _ := reader.ReadString('\n') // Le a nova nota do aluno como texto.
			notaTexto = strings.TrimSpace(notaTexto) // Remove espacos e o \n do fim da nova nota.

			novaNota, err := strconv.ParseFloat(notaTexto, 64) // Converte a nova nota de texto para float64.
			if err != nil { // Verifica erro na conversao da nova nota.
				fmt.Println("Nota inválida. Por favor insira um número!") // Mensagem de erro.
				continue // Volta para tentar novamente.
			}

			if registro.atualizarNotas(nome, novaNota) { // Tenta atualizar a nota do aluno no registro de notas.
				fmt.Println("Nota atualizada com sucesso!") // Mensagem de sucesso se a atualização for bem-sucedida.
			} else {
				fmt.Println("Aluno não encontrado. Não foi possível atualizar a nota.") // Mensagem de erro se o aluno não for encontrado para atualização.
			}

		case "4":
			media := registro.calcularMedia() // Calcula a media das notas dos alunos no registro.
			fmt.Printf("A média das notas dos alunos é: %.2f\n", media) // Exibe a media formatada com duas casas decimais.

		case "5":
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
			fmt.Println("Saindo do programa...") // Mensagem de saida.
			os.Exit(0) // Encerra o programa.

		default: // Caso para opção inválida.
			fmt.Println("Opção inválida. Por favor escolha uma opção entre 1 e 7.") // Mensagem de erro para opção inválida.
		}
	}
	
}