package main

import (
	"fmt"

	"github.com/ffsales/go-trello-poc/config"
)

func main() {
	fmt.Println("Iniciando o |||::Go Trello::|||")

	fmt.Println(config.GetDBConfig().Database)
}
