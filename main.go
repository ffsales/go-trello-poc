package main

import (
	"database/sql"
	"fmt"

	"github.com/ffsales/go-trello-poc/config"
	"github.com/ffsales/go-trello-poc/db"
	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/repository"
)

func main() {
	fmt.Println("Iniciando o |||::Go Trello::|||")

	fmt.Println(config.GetDBConfig().User)

	conn := db.GetConnection()
	defer conn.Close()

	fmt.Println(conn.Ping())

	testInsertCard(conn)

}

func testInsertCard(conn *sql.DB) {
	card := new(models.Card)
	card.Name = "Inicio"
	card.Finished = true

	repository.Insert(conn, *card)
}
