package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoringTimes = 3
const monitoringDealy = 5

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
	fmt.Println("")

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	sites := []string{"http://www.alura.com.br", "http://www.caelum.com.br", "http://www.google.com"}

	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}

		fmt.Println("")
		time.Sleep(monitoringDealy * time.Second)
	}

	fmt.Println("")

}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "Está com problemas. Status code:", resp.StatusCode)
	}
}
