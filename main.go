package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type ipAddress struct {
	ip string
}

func main() {

	menu()

	for {
		inicio()
		switch comando() {
		case 1:
			fmt.Println("")
			fmt.Println("Cadastro de ip")
			cadastroDeIp(stringUser())
		case 2:
			fmt.Println("")
			fmt.Println("Acompanhamento....")
		case 3:
			fmt.Println("")
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("")
			fmt.Println("Saindo da Aplicação...")
			os.Exit(-1)
		}
		time.Sleep(2)
	}
}

func menu() {
	fmt.Println("___________Bem vindo___________")
	fmt.Println()
	fmt.Println("Selecione a Opção desejada")
	fmt.Println()
}
func inicio() {
	fmt.Println("1 - Cadastrar Ip")
	fmt.Println("2 - Acompanhamento em tempo Real")
	fmt.Println("3 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi: ", comando)
	fmt.Println()
	return comando
}

func cadastroDeIp(content string) {
	//os.OpenFile("registro_ip.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	ip, _ := json.Marshal(ipAddress{ip: content})
	fmt.Println(string(ip))
	_ = ioutil.WriteFile("test.json", ip, 0666)

}

func stringUser() string {
	var stringUser string
	fmt.Scan(&stringUser)
	return stringUser
}
