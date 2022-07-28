package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ffsales/go-trello-poc/db"
	"github.com/ffsales/go-trello-poc/repository"
	"github.com/go-chi/chi"
)

type resource struct{}

func GetResource() *resource {
	return new(resource)
}

func (rsc resource) ListBoards(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	boards, _ := repository.GetAllBoards(conn)

	for idx, _ := range boards {
		lists, _ := repository.GetListsByBoard(conn, int(boards[idx].Id))
		for _, list := range lists {
			boards[idx].Lists = append(boards[idx].Lists, &list)
		}
	}

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

	lists, _ := repository.GetListsByBoard(conn, int(board.Id))
	for _, list := range lists {
		board.Lists = append(board.Lists, &list)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(board)
}
