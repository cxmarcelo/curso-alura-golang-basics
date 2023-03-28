package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			printLogs()
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
	sites := readSitesFile()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "Está com problemas. Status code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	//file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)

		if err == io.EOF {
			break
		}

	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()

}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(file))

}
