package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntroduction()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	name := "Marcelo"
	version := 1.2
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	site := "http://www.alura.com.br"

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "Está com problemas. Status code:", resp.StatusCode)

	}
}
