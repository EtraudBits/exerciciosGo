package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Lista struct {
	Items []string
	Quantidade int
}

func (l *Lista) AdicionarItem(item string) { // Metodo para adicionar um item a lista.
	l.Items = append(l.Items, item) // Adiciona o item ao slice de itens.
	l.Quantidade++ // Incrementa a quantidade de itens.
}

func (l *Lista) RemoverItem(item string) { // Metodo para remover um item da lista.
	for i, v := range l.Items { // Loop para encontrar o item a ser removido.
		if v == item { // Verifica se o item atual e o que queremos remover.
			l.Items = append(l.Items[:i], l.Items[i+1:]...) // l.Items[:i] pega os itens antes do indice i e l.Items[i+1:] pega os itens depois do indice i, e os junta para remover o item. e ... e usado para expandir o slice resultante.
			l.Quantidade-- // Decrementa a quantidade de itens.
			return // Sai do metodo apos remover o item.
		}}}

func (l *Lista) ExibirLista() { // Metodo para exibir os itens da lista.
	fmt.Println("Lista de Itens:") // Exibe o titulo da lista.
	for i, item := range l.Items { // Loop para exibir cada item da lista.
		fmt.Printf("%d. %s\n", i+1, item) // Exibe o numero do item e o nome do item.
	}}

func (l *Lista) LimparLista() { // Metodo para limpar a lista.
	l.Items = []string{} // Reseta o slice de itens para um slice vazio.
	l.Quantidade = 0 // Reseta a quantidade de itens para zero.
}

func (l *Lista) Sair() {
	if input := bufio.NewReader(os.Stdin); input != nil { // Verifica se a leitura da entrada padrao e bem sucedida.
		fmt.Println("Saindo da lista...") // Exibe a mensagem de saida.
		os.Exit(0) // Encerra o programa.
	}
	fmt.Println("Saindo da lista...") // Exibe a mensagem de saida.
	os.Exit(0) // Encerra o programa.
}

func main() {
 
	reader := bufio.NewReader(os.Stdin) // Cria um leitor para a entrada padrao

	lista := Lista{} // Cria uma nova instancia da struct Lista.

	for { // Loop principal do programa.
		fmt.Println("\nEscolha uma opcao:") // Exibe as opcoes para o usuario.
		fmt.Println("1. Adicionar item") // Opcao para adicionar um item.
		fmt.Println("2. Remover item") // Opcao para remover um item.
		fmt.Println("3. Exibir lista") // Opcao para exibir a lista.
		fmt.Println("4. Limpar lista") // Opcao para limpar a lista.
		fmt.Println("5. Sair") // Opcao para sair do programa.
		
		opcao, err := reader.ReadString('\n') // Le a opcao escolhida pelo usuario.
		if err != nil { // Verifica erro de leitura.
			fmt.Println("Erro ao ler a opcao. Opções validas são 1, 2, 3, 4 ou 5. Tente novamente!") // Mensagem de erro.
			continue // Volta para tentar novamente.
		}

				opcao = strings.TrimSpace(opcao) // Remove espacos e o \n do fim.

		switch opcao { // Verifica a opcao escolhida usando switch para validar as opcoes.
		case "1": // Caso para adicionar um item.
			fmt.Print("Digite o item a ser adicionado: ") // Prompt para o item a ser adicionado.
			item, _ := reader.ReadString('\n') // Le o item como texto.
			item = strings.TrimSpace(item) // Remove espacos e o \n do fim.
			lista.AdicionarItem(item) // Chama o metodo para adicionar o item a lista.
			fmt.Printf("Item '%s' adicionado!\n", item) // Mensagem de confirmacao.
		case "2": // Caso para remover um item.
			fmt.Print("Digite o item a ser removido: ") // Prompt para o item a ser removido.
			item, _ := reader.ReadString('\n') // Le o item como texto.
			item = strings.TrimSpace(item) // Remove espacos e o \n do fim.
			lista.RemoverItem(item) // Chama o metodo para remover o item da lista.
			fmt.Printf("Item '%s' removido (se existia)!\n", item) // Mensagem de confirmacao.
		case "3": // Caso para exibir a lista.
			lista.ExibirLista() // Chama o metodo para exibir os itens da lista.
		case "4": // Caso para limpar a lista.
			lista.LimparLista() // Chama o metodo para limpar a lista.
			fmt.Println("Lista limpa!") // Mensagem de confirmacao.
		case "5": // Caso para sair do programa.
			lista.Sair() // Chama o metodo para sair do programa.
		default: // Caso para opcao invalida.
			fmt.Println("Opcao invalida. Por favor escolha uma opcao entre 1 e 5.") // Mensagem de erro.
		}
	}

}