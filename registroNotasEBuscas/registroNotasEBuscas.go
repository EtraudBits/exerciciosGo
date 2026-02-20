package main

import (
	"bufio"
	"encoding/json"
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
	indices map[string]int
}

func (r *registroNotas) adicionarAluno(nome string, notas [4]float64) { // Metodo para adicionar um aluno ao registro de notas.
	if r.indices == nil {
		r.indices = make(map[string]int)
	}
	indice := len(r.alunos)
	r.alunos = append(r.alunos, aluno{nome: nome, notas: notas}) // Adiciona um novo aluno ao slice de alunos do registro.
	r.indices[nome] = indice // Atualiza o mapa de índices para permitir busca rápida por nome.
}

func (r *registroNotas) salvar() error { // Metodo para salvar o registro de notas em um arquivo JSON.
	dados, err := json.MarshalIndent(r.alunos, "", " ") // Converte o slice de alunos para JSON formatado.
	if err != nil { // Verifica se houve um erro na conversão para JSON.
		return err // Retorna o erro se a conversão falhar.
	}
	return os.WriteFile("notas.json", dados, 0644) // Escreve os dados JSON em um arquivo chamado "notas.json" com permissões de leitura e escrita.
}

func (r *registroNotas) buscarAluno(nome string) *aluno { // Metodo para buscar um aluno pelo nome no registro de notas.
	if r.indices == nil {
		return nil
	}
	indice, ok := r.indices[nome] // Busca o índice do aluno no mapa de índices em tempo constante O(1).
	if !ok {
		return nil // Retorna nil se o aluno não for encontrado.
	}
	return &r.alunos[indice] // Retorna o ponteiro para o aluno encontrado.
}

func (r *registroNotas) atualizarNotas(nome string, novasNotas [4]float64) bool { // Metodo para atualizar as notas de um aluno no registro de notas.
	if r.indices == nil {
		return false
	}
	indice, ok := r.indices[nome]
	if !ok {
		return false // Retorna false se o aluno não for encontrado.
	}
	r.alunos[indice].notas = novasNotas // Atualiza as notas do aluno encontrado.
	return true // Retorna true para indicar que a atualização foi bem-sucedida.
}

func (r *registroNotas) incluirNotaAluno(nome string, numeroNota int, nota float64) bool { // Metodo para incluir uma nota em um aluno existente no registro de notas, especificando qual nota (1 a 4) deve ser atualizada.
	if numeroNota < 1 || numeroNota > 4 { // Verifica se o numero da nota é válido (entre 1 e 4). Se não for, retorna false para indicar que a inclusão da nota falhou.
		return false // Retorna false para indicar que a inclusão da nota falhou devido a um número de nota inválido.
	}

	if r.indices == nil {
		return false
	}
	indice, ok := r.indices[nome]
	if !ok {
		return false // Retorna false se o aluno não for encontrado para incluir a nota.
	}
	r.alunos[indice].notas[numeroNota-1] = nota // Atualiza a nota específica do aluno encontrado.
	return true // Retorna true para indicar que a inclusão da nota foi bem-sucedida.
}

func (r *registroNotas) calcularMediaAluno(nome string) (float64, bool) { // Metodo para calcular a media de um aluno no registro de notas.
	if r.indices == nil {
		return 0.0, false
	}
	indice, ok := r.indices[nome]
	if !ok {
		return 0.0, false // Retorna 0.0 e false se o aluno não for encontrado para calcular a media.
	}
	a := r.alunos[indice]
	soma := 0.0 // Variavel para armazenar a soma das notas do aluno.
	for _, nota := range a.notas { // Loop para somar as notas do aluno.
		soma += nota // Adiciona a nota atual à soma total das notas do aluno.
	}
	return soma / 4.0, true // Retorna a media calculada (soma das notas dividida por 4) e true para indicar que o cálculo foi bem-sucedido.
}

func (r *registroNotas) removerAluno(nome string) bool { // Metodo para remover um aluno do registro de notas.
	if r.indices == nil {
		return false
	}
	indice, ok := r.indices[nome]
	if !ok {
		return false // Retorna false se o aluno não for encontrado.
	}

	r.alunos = append(r.alunos[:indice], r.alunos[indice+1:]...) // Remove o aluno do slice de alunos usando slicing.
	delete(r.indices, nome) // Remove o aluno do mapa de índices.

	// Atualizar índices de todos os alunos após o índice removido
	for i := indice; i < len(r.alunos); i++ {
		r.indices[r.alunos[i].nome] = i // Reajusta os índices no mapa após a remoção.
	}

	return true // Retorna true para indicar que a remoção foi bem-sucedida.
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

func lerNota(reader *bufio.Reader, mensagem string) (float64, bool) { // Funcao para ler uma nota do usuario, com validacao de entrada.
	fmt.Print(mensagem) // Exibe a mensagem de prompt para o usuario, solicitando a nota.
	notaTexto, err := reader.ReadString('\n') // Le a entrada do usuario como uma string.
	if err != nil { // Verifica se houve um erro ao ler a entrada do usuario.
		fmt.Println("Erro ao ler a nota.") // Mensagem de erro se houver um problema ao ler a nota.
		return 0, false // Retorna 0 e false para indicar que a leitura da nota falhou.
	}

	notaTexto = strings.TrimSpace(notaTexto) // Remove espacos e o \n do fim da string da nota, para garantir que a conversao funcione corretamente.
	nota, err := strconv.ParseFloat(notaTexto, 64) // Tenta converter a string da nota para um valor float64. Se a conversao falhar, o erro sera tratado no bloco if abaixo.
	if err != nil { // Verifica se houve um erro na conversao da nota.
		fmt.Println("Nota inválida. A nota deve ser um número.") // Mensagem de erro se a nota nao for um numero valido.
		return 0, false // Retorna 0 e false para indicar que a conversao da nota falhou.
	}

	return nota, true // Retorna a nota convertida e true para indicar que a leitura e conversao da nota foram bem-sucedidas.
}

func nomeValido(nome string) bool { // Funcao para validar o nome do aluno, garantindo que seja composto apenas por letras e espacos, e que nao seja vazio.
	nome = strings.TrimSpace(nome) // Remove espacos em branco do inicio e fim do nome.
	if nome == "" { // Verifica se o nome é vazio após remover os espaços. Se for vazio, retorna false para indicar que o nome é inválido.
		return false //	 Retorna false para indicar que o nome é inválido.
	}

	for _, r := range nome { // Loop para verificar cada caractere do nome, garantindo que seja uma letra ou um espaço. Se encontrar um caractere inválido, retorna false.
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) { // Verifica se o caractere não é uma letra nem um espaço.
			return false // Retorna false para indicar que o nome é inválido se encontrar um caractere que não seja letra ou espaço.
		}
	}

	return true // Retorna true para indicar que o nome é válido.
}

func confirmarAcao(reader *bufio.Reader, pergunta string) bool { // Funcao para confirmar uma acao do usuario, retornando true se a resposta for afirmativa.
	fmt.Print(pergunta) // Exibe a pergunta para o usuario.
	resposta, err := reader.ReadString('\n') // Le a resposta do usuario como uma string.
	if err != nil { // Verifica se houve um erro ao ler a resposta.
		fmt.Println("Erro ao ler confirmação. Operação cancelada.") // Mensagem de erro se houver um problema ao ler a resposta.
		return false // Retorna false para indicar que a confirmacao falhou.
	}

	resposta = strings.TrimSpace(strings.ToLower(resposta)) // Remove espacos em branco e converte a resposta para minusculas.
	return resposta == "s" || resposta == "sim" // Retorna true se a resposta for "s" ou "sim".
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

			var notas [4]float64 // Variavel para armazenar as 4 notas do aluno.
			valido := true // Variavel para controlar se as notas foram lidas com sucesso.
			for i := 0; i < 4; i++ { // Loop para ler as 4 notas do aluno, usando a funcao lerNota para validar cada nota individualmente.
				nota, ok := lerNota(reader, fmt.Sprintf("Digite a nota %d do aluno: ", i+1)) // Prompt para a nota do aluno, indicando qual numero da nota (1 a 4) está sendo solicitada.
				if !ok { // Verifica se a leitura da nota foi bem-sucedida. Se não for, define valido como false e quebra o loop para evitar adicionar um aluno com notas inválidas.
					valido = false // Define valido como false para indicar que houve um problema na leitura de uma das notas.
					break // Sai do loop de leitura das notas, pois não faz sentido continuar se uma nota for inválida.
				}
				notas[i] = nota // Armazena a nota lida na posição correta do array de notas do aluno.
			}

			if !valido { // Verifica se todas as notas foram lidas com sucesso. Se não foram, exibe uma mensagem de erro e continua para a próxima iteração do loop principal, evitando adicionar um aluno com notas inválidas.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente adicionar um aluno novamente com notas válidas.
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

			var notas [4]float64 // Variavel para armazenar as novas 4 notas do aluno.
			valido := true // Variavel para controlar se as novas notas foram lidas com sucesso.
			for i := 0; i < 4; i++ { // Loop para ler as novas 4 notas do aluno, usando a funcao lerNota para validar cada nota individualmente.
				nota, ok := lerNota(reader, fmt.Sprintf("Digite a nova nota %d do aluno: ", i+1)) // Prompt para a nova nota do aluno, indicando qual numero da nota (1 a 4) está sendo solicitada.
				if !ok { // Verifica se a leitura da nova nota foi bem-sucedida. Se não for, define valido como false e quebra o loop para evitar atualizar um aluno com notas inválidas.
					valido = false // Define valido como false para indicar que houve um problema na leitura de uma das novas notas.
					break // Sai do loop de leitura das novas notas, pois não faz sentido continuar se uma nova nota for inválida.
				}
				notas[i] = nota // Armazena a nova nota lida na posição correta do array de novas notas do aluno.
			}

			if !valido { 	// Verifica se todas as novas notas foram lidas com sucesso. Se não foram, exibe uma mensagem de erro e continua para a próxima iteração do loop principal, evitando atualizar um aluno com notas inválidas.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente atualizar um aluno novamente com notas válidas.
			}

			if registro.atualizarNotas(nome, notas) { // Tenta atualizar as notas do aluno no registro de notas.
				fmt.Println("Nota atualizada com sucesso!") // Mensagem de sucesso se a atualização for bem-sucedida.
			} else {
				fmt.Println("Aluno não encontrado. Não foi possível atualizar a nota.") // Mensagem de erro se o aluno não for encontrado para atualização.
			}

		case "4":
			fmt.Print("Digite o nome do aluno para calcular a média: ") // Prompt para o nome do aluno para calcular a media.
			nome, _ := reader.ReadString('\n') // Le o nome do aluno como texto.
			nome = strings.TrimSpace(nome) // Remove espacos e o \n do fim do nome.
			if !nomeValido(nome) { // Verifica se o nome do aluno é válido.
				fmt.Println("Aluno inválido. Digite um nome com texto (apenas letras e espaços).") // Mensagem de erro se o nome do aluno for inválido.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente calcular a média de um aluno novamente com um nome válido.
			}

			media, ok := registro.calcularMediaAluno(nome) // Tenta calcular a media do aluno no registro de notas.
			if !ok { // Verifica se o cálculo da média foi bem-sucedido. Se não for, exibe uma mensagem de erro indicando que o aluno não foi encontrado.
				fmt.Println("Aluno não encontrado.") // Mensagem de erro se o aluno não for encontrado para calcular a média.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente calcular a média de um aluno novamente.
			}

			fmt.Printf("A média do aluno %s é: %.2f\n", nome, media) 	// Exibe a média calculada do aluno, formatada com 2 casas decimais.

		case "5":
			if !confirmarAcao(reader, "Deseja remover um aluno? (s/n): ") { // Confirmação para remover um aluno, perguntando ao usuário se ele tem certeza de que deseja realizar a remoção.
				if !confirmarAcao(reader, "Tem certeza que deseja remover um aluno? Esta ação não pode ser desfeita. (s/n): ") {
					fmt.Println("Remoção cancelada.") // Mensagem de cancelamento se o usuário não confirmar a remoção.
					continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente remover um aluno novamente.
				}
				fmt.Println("Remoção cancelada.") // Mensagem de cancelamento se o usuário não confirmar a remoção.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente remover um aluno novamente.
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
			fmt.Print("Digite o nome do aluno para incluir nova nota: ") // Prompt para o nome do aluno para incluir uma nova nota.
			nome, _ := reader.ReadString('\n') // Le o nome do aluno como texto.
			nome = strings.TrimSpace(nome) // Remove espacos e o \n do fim do nome.
			if !nomeValido(nome) { // Verifica se o nome do aluno é válido.
				fmt.Println("Aluno inválido. Digite um nome com texto (apenas letras e espaços).") // Mensagem de erro se o nome do aluno for inválido.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente incluir uma nova nota novamente com um nome válido.
			}

			fmt.Print("Escolha qual nota deseja incluir (1, 2, 3 ou 4): ") // Prompt para o número da nota a ser incluída, indicando que deve ser um número entre 1 e 4.
			numeroTexto, _ := reader.ReadString('\n') // Le o número da nota como texto.
			numeroTexto = strings.TrimSpace(numeroTexto) // Remove espacos e o \n do fim do número da nota para garantir que a conversão para inteiro funcione corretamente.

			numeroNota, err := strconv.Atoi(numeroTexto) // Tenta converter o número da nota de string para inteiro. Se a conversão falhar, o erro será tratado no bloco if abaixo.
			if err != nil || numeroNota < 1 || numeroNota > 4 { // Verifica se houve um erro na conversão do número da nota ou se o número da nota não está entre 1 e 4. Se qualquer uma dessas condições for verdadeira, exibe uma mensagem de erro e continua para a próxima iteração do loop principal, permitindo que o usuário tente incluir uma nova nota novamente com um número de nota válido.
				fmt.Println("Opção inválida. Escolha um número entre 1 e 4.") 	// Mensagem de erro se o número da nota for inválido.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente incluir uma nova nota novamente com um número de nota válido.
			}

			novaNota, ok := lerNota(reader, "Digite o valor da nota: ") // Prompt para o valor da nova nota a ser incluída, usando a função lerNota para validar a entrada do usuário. Se a leitura da nota falhar, o erro será tratado no bloco if abaixo.
			if !ok { // Verifica se a leitura da nova nota foi bem-sucedida. Se não for, exibe uma mensagem de erro e continua para a próxima iteração do loop principal, permitindo que o usuário tente incluir uma nova nota novamente com um valor de nota válido.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário tente incluir uma nova nota novamente com um valor de nota válido.
			}

			if registro.incluirNotaAluno(nome, numeroNota, novaNota) { 	// Tenta incluir a nova nota para o aluno no registro de notas, especificando o número da nota a ser atualizada. Se a inclusão for bem-sucedida, exibe uma mensagem de sucesso. Caso contrário, exibe uma mensagem de erro indicando que o aluno não foi encontrado para incluir a nota.
				fmt.Println("Nota incluída com sucesso!")
			} else {
				fmt.Println("Aluno não encontrado. Não foi possível incluir a nota.")
			}

		case "8":
			if !confirmarAcao(reader, "Deseja sair do programa? (s/n): ") { // Prompt para confirmar se o usuário deseja sair do programa.
				fmt.Println("Saída cancelada.") // Mensagem de cancelamento de saída.
				continue // Continua para a próxima iteração do loop principal, permitindo que o usuário continue usando o programa.
			}

			fmt.Println("Saindo do programa...") // Mensagem de saida.
			os.Exit(0) // Encerra o programa.

		default: // Caso para opção inválida.
			fmt.Println("Opção inválida. Por favor escolha uma opção entre 1 e 8.") // Mensagem de erro para opção inválida.
		}
	}
	
}