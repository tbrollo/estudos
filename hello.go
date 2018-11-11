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

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do Programa...")
		// } else {
		// 	fmt.Println("Não conheço este comando.")
		// }

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimirLogs()
		case 9:
			fmt.Println("Saindo do Programa...")
			os.Exit(0) //saiu com sucesso
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1) //ocorreu algum erro
		}
		fmt.Println()
	}
}

func exibeIntroducao() {
	nome := "Tabata"
	versao := 1.1
	fmt.Println("Olá, sra.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println()
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("9- Sair do Programa")
}

func leComando() int {
	fmt.Print(">> ")

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println()
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// var sites [4]string
	// sites[0] = "https://random-status-code.herokuapp.com/"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.caelum.com.br"

	//sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		println()
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}
func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Ocorreu um erro:", err)
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimirLogs() {
	fmt.Println("Exibindo Logs...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
	}

	fmt.Println(string(arquivo))

}
