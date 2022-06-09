package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/wesleymr59/PingPong-Golang/models"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type IpAddress struct {
	ip   int
	nome string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("teste banco postgres")
	initDataBase()
	fmt.Println()
	menu()

	for {
		inicio()
		switch comando() {
		case 1:
			fmt.Println("")
			fmt.Println("Cadastro de ip")
			INSERTIpDataBase(stringUser())
		case 2:
			fmt.Println("")
			fmt.Println("Acompanhamento....")
		case 3:
			fmt.Println("")
			fmt.Println("Exibindo Logs...")
		case 4:
			fmt.Println("")
			fmt.Println("Exibindo Lista de Ips...")
			SELECTIpDataBase()
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
	fmt.Println("4 - Buscar Ip's")
	fmt.Println("0 - Sair do Programa")
}

func comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi: ", comando)
	fmt.Println()
	return comando
}

//func cadastroDeIp(content string) {
//	//os.OpenFile("registro_ip.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	ip, _ := json.Marshal(IpAddress{ip: content})
//	fmt.Println(string(ip))
//	_ = ioutil.WriteFile("test.json", ip, 0666)
//}

func stringUser() (int, string) {
	var ip int
	var nome string
	fmt.Scan(&ip)
	fmt.Scan(&nome)
	return ip, nome
}

func connPosrgre() string {
	var stringConn string = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRE_HOST"), os.Getenv("POSTGRE_PORT"), os.Getenv("POSTGRE_USER"), os.Getenv("POSTGRE_PASSWORD"), os.Getenv("POSTGRE_DBNAME"))

	fmt.Println("print vagabundo")
	fmt.Println(stringConn)
	return stringConn
}

func initDataBase() {

	db, err := sql.Open("postgres", connPosrgre())

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado Pae!")
}

func INSERTIpDataBase(ip int, nome string) {

	db, err := sql.Open("postgres", connPosrgre())

	sqlInsert := `INSERT INTO ip (ip, nome)
	VALUES ($1, $2)`
	_, err = db.Exec(sqlInsert, ip, nome)

	if err != nil {
		panic(err)
	}

	db.Close()
}

func SELECTIpDataBase() {

	db, err := sql.Open("postgres", connPosrgre())

	sqlSelect := `SELECT * FROM IP`

	var ipAddress IpAddress
	row := db.QueryRow(sqlSelect, 2)
	err = row.Scan(&ipAddress.ip, &ipAddress.nome)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Não existe colunas para retorno!")
		return
	case nil:
		fmt.Println("user")
	default:
		panic(err)
	}
}
