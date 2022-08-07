package cards

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ffsales/go-trello-poc/db"
	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetCardsByList(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strListId := chi.URLParam(r, "listId")
	listId, err := strconv.Atoi(strListId)
	if err != nil {
		panic(err)
	}

	cards, _ := repository.GetCardsByList(conn, listId)

	w.Header().Set("Content-Type", "application/json")

	respCards := []render.Renderer{}

	for _, card := range cards {
		respCards = append(respCards, card.ToResponse())
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, respCards)
}

func GetAllCards(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	cards, _ := repository.GetAllCards(conn)

	w.Header().Set("Content-Type", "application/json")

	respCards := []render.Renderer{}

	for _, card := range cards {
		respCards = append(respCards, card.ToResponse())
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, respCards)
}

func GetCard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strCardId := chi.URLParam(r, "cardId")
	cardId, err := strconv.Atoi(strCardId)
	if err != nil {
		panic(err)
	}

	card, _ := repository.GetCard(conn, cardId)

	w.Header().Set("Content-Type", "application/json")

	render.Status(r, http.StatusOK)
	render.Render(w, r, card.ToResponse())
}

func CreateCard(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		panic("Body empty!")
	}

	conn := db.GetConnection()
	defer conn.Close()

	var requestCard models.Card
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestCard); err != nil {
		panic(err)
	}

	if requestCard.IdList <= 0 {
		panic("erro")
	}

	if list, err := repository.GetList(conn, int(requestCard.IdList)); err != nil || list.Id <= 0 {
		panic(err)
	}

	card, err := repository.InsertCard(conn, requestCard)
	if err != nil {
		panic(err)
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, card.ToResponse())
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		panic("Body empty!")
	}

	conn := db.GetConnection()
	defer conn.Close()

	strCardId := chi.URLParam(r, "cardId")
	cardId, err := strconv.Atoi(strCardId)
	if err != nil {
		panic(err)
	}

	var requestCard models.Card
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestCard); err != nil {
		panic(err)
	}

	foundCard, err := repository.GetCard(conn, int(cardId))
	if err != nil {
		panic(err)
	}

	foundCard.Name = requestCard.Name
	foundCard.Finished = requestCard.Finished

	if rows, err := repository.UpdateCard(conn, &foundCard); err != nil {
		panic(err)
	} else if rows != 1 {
		panic(fmt.Sprintf("Error: %d rows affected", rows))
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, foundCard.ToResponse())
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strCardId := chi.URLParam(r, "cardId")
	cardId, err := strconv.Atoi(strCardId)
	if err != nil {
		panic(err)
	}

	if rows, err := repository.DeleteCard(conn, cardId); err != nil {
		panic(err)
	} else if rows != 1 {
		panic(fmt.Sprintf("Error: %d rows affected", rows))
	}

	w.WriteHeader(http.StatusNoContent)
}
