package lists

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ffsales/go-trello-poc/db"
	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/repository"
	"github.com/go-chi/chi"
)

func GetListsByBoard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strBoardId := chi.URLParam(r, "boardId")
	boardId, err := strconv.Atoi(strBoardId)
	if err != nil {
		panic(err)
	}

	lists, _ := repository.GetListsByBoard(conn, boardId)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(lists)
}

func GetAllLists(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	lists, _ := repository.GetAllLists(conn)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(lists)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strListId := chi.URLParam(r, "listId")
	listId, err := strconv.Atoi(strListId)
	if err != nil {
		panic(err)
	}

	list, _ := repository.GetList(conn, listId)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(list)
}

func CreateList(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		panic("Body empty!")
	}

	conn := db.GetConnection()
	defer conn.Close()

	var requestList models.List
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestList); err != nil {
		panic(err)
	}

	if requestList.IdBoard <= 0 {
		panic("erro")
	}

	if board, err := repository.GetBoard(conn, requestList.IdBoard); err != nil || board.Id <= 0 {
		panic("board inválido")
	}

	respList, err := repository.InsertList(conn, requestList)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respList)
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		panic("Body empty!")
	}

	conn := db.GetConnection()
	defer conn.Close()

	strListId := chi.URLParam(r, "listId")
	listId, err := strconv.Atoi(strListId)
	if err != nil {
		panic(err)
	}

	var requestList models.List
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestList); err != nil {
		panic(err)
	}

	foundList, err := repository.GetList(conn, int(listId))
	if err != nil {
		panic(err)
	}

	foundList.Name = requestList.Name
	foundList.Order = requestList.Order

	if rows, err := repository.UpdateList(conn, &foundList); err != nil {
		panic(err)
	} else if rows != 1 {
		panic(fmt.Sprintf("Error: %d rows affected", rows))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundList)
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strListId := chi.URLParam(r, "listId")
	listId, err := strconv.Atoi(strListId)
	if err != nil {
		panic(err)
	}

	if rows, err := repository.DeleteList(conn, listId); err != nil {
		panic(err)
	} else if rows != 1 {
		panic(fmt.Sprintf("Error: %d rows affected", rows))
	}

	w.WriteHeader(http.StatusNoContent)
}
