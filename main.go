package main

import (
	"database/sql"
	"fmt"
	"log"

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

	// fmt.Println(conn.Ping())

	// testInsertCard(conn)
	// testGetCard(conn)
	// testGetAllCards(conn)
	// testUpdateCard(conn)
	// testDeleteCard(conn)
}

func testInsertCard(conn *sql.DB) {
	var card models.Card = models.Card{
		Name:     "Teste",
		Finished: true,
	}

	newCard, err := repository.Insert(conn, card)
	fmt.Println(err)
	fmt.Println(newCard)
}

func testGetCard(conn *sql.DB) {

	card, err := repository.Get(conn, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(card)
}

func testGetAllCards(conn *sql.DB) {
	cards, err := repository.GetAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cards)
}

func testUpdateCard(conn *sql.DB) {
	card := new(models.Card)
	card.Id = 1
	card.Name = "Meio"
	card.Finished = true

	rows, _ := repository.Update(conn, card)
	fmt.Println(rows)

	newCard, _ := repository.Get(conn, int(card.Id))
	fmt.Println(newCard)
}

func testDeleteCard(conn *sql.DB) {

	fmt.Println(repository.GetAll(conn))

	repository.Delete(conn, 7)

	fmt.Println(repository.GetAll(conn))
}
