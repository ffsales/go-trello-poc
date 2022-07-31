package board

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

type resource struct{}

func GetResource() *resource {
	return new(resource)
}

func (rsc resource) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	boards, _ := repository.GetAllBoards(conn)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(boards)
}

func (rsc resource) GetBoard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strBoardId := chi.URLParam(r, "boardId")
	boardId, err := strconv.Atoi(strBoardId)
	if err != nil {
		panic(err)
	}

	board, err := repository.GetBoard(conn, boardId)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(board)
}

func (rsc resource) CreateBoard(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		panic("Body empty!")
	}

	conn := db.GetConnection()
	defer conn.Close()

	var requestBoard models.Board
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBoard); err != nil {
		panic(err)
	}

	respBoard, err := repository.InsertBoard(conn, requestBoard)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respBoard)
}

func (rsc resource) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		panic("Body empty!")
	}

	conn := db.GetConnection()
	defer conn.Close()

	strBoardId := chi.URLParam(r, "boardId")
	boardId, err := strconv.Atoi(strBoardId)
	if err != nil {
		panic(err)
	}

	var requestBoard models.Board
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBoard); err != nil {
		panic(err)
	}

	foundBoard, err := repository.GetBoard(conn, int(boardId))
	if err != nil {
		panic(err)
	}

	foundBoard.Name = requestBoard.Name
	foundBoard.Description = requestBoard.Description

	if rows, err := repository.UpdateBoard(conn, &foundBoard); err != nil {
		panic(err)
	} else if rows != 1 {
		panic(fmt.Sprintf("Error: %d rows affected", rows))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(foundBoard)
}

func (rsc resource) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strBoardId := chi.URLParam(r, "boardId")
	boardId, err := strconv.Atoi(strBoardId)
	if err != nil {
		panic(err)
	}

	if rows, err := repository.DeleteBoard(conn, boardId); err != nil {
		panic(err)
	} else if rows != 1 {
		panic(fmt.Sprintf("Error: %d rows affected", rows))
	}

	w.WriteHeader(http.StatusNoContent)
}
