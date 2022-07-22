package main

import (
	"fmt"

	"github.com/ffsales/go-trello-poc/config"
	"github.com/ffsales/go-trello-poc/db"
)

func main() {
	fmt.Println("Iniciando o |||::Go Trello::|||")

	fmt.Println(config.GetDBConfig().User)

	conn := db.GetConnection()
	defer conn.Close()

	fmt.Println(conn.Ping())

}
