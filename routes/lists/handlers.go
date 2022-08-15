package lists

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

func GetListsByBoard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strBoardId := chi.URLParam(r, "boardId")
	boardId, err := strconv.Atoi(strBoardId)
	if err != nil {
		utils.BadRequestError(w, r, err, "Invalid id")
		return
	}

	lists, err := repository.GetListsByBoard(conn, boardId)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	}

	respLists := []render.Renderer{}

	for _, list := range lists {
		respLists = append(respLists, list.ToResponse())
	}

	w.Header().Set("Content-Type", "application/json")

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, respLists)
}

func GetAllLists(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	lists, err := repository.GetAllLists(conn)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	}

	respLists := []render.Renderer{}

	for _, list := range lists {
		respLists = append(respLists, list.ToResponse())
	}

	w.Header().Set("Content-Type", "application/json")

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, respLists)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strListId := chi.URLParam(r, "listId")
	listId, err := strconv.Atoi(strListId)
	if err != nil {
		utils.BadRequestError(w, r, err, "Invalid error")
		return
	}

	list, _ := repository.GetList(conn, listId)

	w.Header().Set("Content-Type", "application/json")

	render.Status(r, http.StatusOK)
	render.Render(w, r, list.ToResponse())
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Body == nil {
		panic("Body empty!")
	}

	conn := db.GetConnection()
	defer conn.Close()

	var requestList models.List
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestList); err != nil {
		utils.UnprocessableEntityError(w, r, err, "Invalid body")
		return
	}

	if requestList.IdBoard <= 0 {
		utils.UnprocessableEntityError(w, r, err, "Invalid body")
		return
	}

	if board, err := repository.GetBoard(conn, requestList.IdBoard); err != nil || board.Id <= 0 {
		utils.NotFoundError(w, r, err, "Not found")
		return
	}

	list, err := repository.InsertList(conn, requestList)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, list.ToResponse())
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Body == nil {
		utils.UnprocessableEntityError(w, r, err, "Empty body")
		return
	}

	conn := db.GetConnection()
	defer conn.Close()

	strListId := chi.URLParam(r, "listId")
	listId, err := strconv.Atoi(strListId)
	if err != nil {
		utils.BadRequestError(w, r, err, "Invalid id")
		return
	}

	var requestList models.List
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestList); err != nil {
		utils.UnprocessableEntityError(w, r, err, "Invalid body")
		return
	}

	foundList, err := repository.GetList(conn, int(listId))
	if err != nil {
		utils.NotFoundError(w, r, err, "Not found")
		return
	}

	foundList.Name = requestList.Name
	foundList.Order = requestList.Order

	if rows, err := repository.UpdateList(conn, &foundList); err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	} else if rows != 1 {
		utils.UnprocessableEntityError(w, r, err, fmt.Sprintf("Error: %d rows affected", rows))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, foundList.ToResponse())
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strListId := chi.URLParam(r, "listId")
	listId, err := strconv.Atoi(strListId)
	if err != nil {
		utils.BadRequestError(w, r, err, "Invalid id")
		return
	}

	if rows, err := repository.DeleteList(conn, listId); err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern error")
		return
	} else if rows != 1 {
		utils.ServiceUnavailableError(w, r, err, fmt.Sprintf("Error: %d rows affected", rows))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
