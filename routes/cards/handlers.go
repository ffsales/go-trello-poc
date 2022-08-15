package cards

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ffsales/go-trello-poc/db"
	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/repository"
	"github.com/ffsales/go-trello-poc/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetCardsByList(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strListId := chi.URLParam(r, "listId")
	listId, err := strconv.Atoi(strListId)
	if err != nil {
		utils.BadRequestError(w, r, err, "Invalid Parameter")
	}

	cards, err := repository.GetCardsByList(conn, listId)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Error Service")
	}

	respCards := []render.Renderer{}

	for _, card := range cards {
		respCards = append(respCards, card.ToResponse())
	}

	utils.OkList(w, r, respCards)
}

func GetAllCards(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	cards, err := repository.GetAllCards(conn)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Error Service")
	}

	respCards := []render.Renderer{}

	for _, card := range cards {
		respCards = append(respCards, card.ToResponse())
	}

	utils.OkList(w, r, respCards)
}

func GetCard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strCardId := chi.URLParam(r, "cardId")
	cardId, err := strconv.Atoi(strCardId)
	if err != nil {
		utils.UnprocessableEntityError(w, r, err, "Invalid Parameter")
		return
	}

	card, err := repository.GetCard(conn, cardId)
	if err != nil || card == (models.Card{}) {
		utils.NotFoundError(w, r, err, "Card Not Found")
		return
	}

	utils.Ok(w, r, card.ToResponse())
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Body == nil {
		utils.UnprocessableEntityError(w, r, err, "Empty body")
		return
	}

	conn := db.GetConnection()
	defer conn.Close()

	var requestCard models.Card
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestCard); err != nil {
		utils.UnprocessableEntityError(w, r, err, "Invalid request")
		return
	}

	if requestCard.IdList <= 0 {
		utils.UnprocessableEntityError(w, r, err, "Invalid request")
		return
	}

	if list, err := repository.GetList(conn, int(requestCard.IdList)); err != nil || list.Id <= 0 {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	}

	card, err := repository.InsertCard(conn, requestCard)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, card.ToResponse())
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Body == nil {
		utils.UnprocessableEntityError(w, r, err, "Empty body")
		return
	}

	conn := db.GetConnection()
	defer conn.Close()

	strCardId := chi.URLParam(r, "cardId")
	cardId, err := strconv.Atoi(strCardId)
	if err != nil {
		utils.BadRequestError(w, r, err, "Invalid id")
		return
	}

	var requestCard models.Card
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestCard); err != nil {
		utils.UnprocessableEntityError(w, r, err, "Invalid body")
		return
	}

	foundCard, err := repository.GetCard(conn, int(cardId))
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	}

	foundCard.Name = requestCard.Name
	foundCard.Finished = requestCard.Finished

	if rows, err := repository.UpdateCard(conn, &foundCard); err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	} else if rows != 1 {
		utils.BadRequestError(w, r, err, fmt.Sprintf("Error: %d rows affected", rows))
		return
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
		utils.BadRequestError(w, r, err, "Invalid id")
		return
	}

	if rows, err := repository.DeleteCard(conn, cardId); err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	} else if rows != 1 {
		utils.BadRequestError(w, r, err, fmt.Sprintf("Error: %d rows affected", rows))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
