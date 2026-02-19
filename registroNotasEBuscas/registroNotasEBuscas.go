package main

import (
	"fmt"
)


type aluno struct {
	nome string
	nota float64
}

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

}