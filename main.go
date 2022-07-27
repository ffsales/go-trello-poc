package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/ffsales/go-trello-poc/config"
	"github.com/ffsales/go-trello-poc/db"
	"github.com/ffsales/go-trello-poc/handlers"
	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("Iniciando o |||::Go Trello::|||")

	fmt.Println(config.GetDBConfig().User)

	conn := db.GetConnection()
	defer conn.Close()

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	boardResource := handlers.GetResource()

	router.Get("/go-trello/boards", boardResource.ListBoards)

	http.ListenAndServe(":3000", router)

	// fmt.Println(conn.Ping())

	// testInsertBoard(conn)
	// testDeleteBoard(conn)

	// testInsertList(conn)
	// testGetList(conn)
	// testGetAllLists(conn)
	// testDeleteList(conn)

	// testInsertCard(conn)
	// testGetCard(conn)
	// testGetAllCards(conn)
	// testUpdateCard(conn)
	// testDeleteCard(conn)
}

func FirstTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func testInsertBoard(conn *sql.DB) {
	var board models.Board = models.Board{
		Name:        "Teste",
		Description: "Board de teste",
	}

	newBoard, err := repository.InsertBoard(conn, board)
	fmt.Println(err)
	fmt.Println(newBoard)
}

func testDeleteBoard(conn *sql.DB) {

	fmt.Println(repository.GetAllBoards(conn))

	repository.DeleteBoard(conn, 2)

	fmt.Println(repository.GetAllBoards(conn))
}

func testInsertList(conn *sql.DB) {
	var list models.List = models.List{
		Name:    "Teste",
		Order:   1,
		IdBoard: 3,
	}

	newList, err := repository.InsertList(conn, list)
	fmt.Println(err)
	fmt.Println(newList)
}

func testGetList(conn *sql.DB) {

	list, err := repository.GetList(conn, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(list)
}

func testGetAllLists(conn *sql.DB) {
	lists, err := repository.GetAllLists(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lists)
}

func testDeleteList(conn *sql.DB) {

	fmt.Println(repository.GetAllLists(conn))

	repository.DeleteList(conn, 5)

	fmt.Println(repository.GetAllLists(conn))
}

func testInsertCard(conn *sql.DB) {
	var card models.Card = models.Card{
		Name:     "Teste",
		Finished: true,
		IdList:   5,
	}

	newCard, err := repository.InsertCard(conn, card)
	fmt.Println(err)
	fmt.Println(newCard)
}

func testGetCard(conn *sql.DB) {

	card, err := repository.GetCard(conn, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(card)
}

func testGetAllCards(conn *sql.DB) {
	cards, err := repository.GetAllCards(conn)
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

	rows, _ := repository.UpdateCard(conn, card)
	fmt.Println(rows)

	newCard, _ := repository.GetCard(conn, int(card.Id))
	fmt.Println(newCard)
}

func testDeleteCard(conn *sql.DB) {

	fmt.Println(repository.GetAllCards(conn))

	repository.DeleteCard(conn, 7)

	fmt.Println(repository.GetAllCards(conn))
}
