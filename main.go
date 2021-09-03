package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var nick string
	fmt.Print("Digite seu Nick: ")
	fmt.Scan(&nick)
	saudacao(nick)

	for {
		menu()
		comando := lerComando()

		switch comando {
		case 1:
			// random-status-code.herokuapp.com
			var urlSite string
			fmt.Print("Digite a url do site(não precisa do http): ")
			fmt.Scan(&urlSite)
			monitoramento(urlSite)
		case 2:
			fmt.Println("Exibindo os Logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido")
			os.Exit(-1)
		}
	}

}

func saudacao(userName string) {
	nome := userName
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Esse programa está na versão:", versao)
}

func menu() {
	fmt.Println()
	fmt.Println("****************************")
	fmt.Println("* 1- Iniciar Monitoramento *")
	fmt.Println("* 2- Exibir Logs           *")
	fmt.Println("* 0- Sair do Programa      *")
	fmt.Println("****************************")
}

func lerComando() int {
	var newComando int
	fmt.Scan(&newComando)

	return newComando
}

func monitoramento(urlSite string) {
	fmt.Println("Buscando...")
	site := "http://" + urlSite
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
		fmt.Println("Status:", response.StatusCode)
	} else {
		fmt.Println("O site:", site, "apresentou problemas ao carregar...")
		fmt.Println("Status:", response.StatusCode)
	}
}
