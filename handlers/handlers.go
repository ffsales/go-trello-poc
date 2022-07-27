package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ffsales/go-trello-poc/db"
	"github.com/ffsales/go-trello-poc/repository"
)

type resource struct{}

func GetResource() *resource {
	return new(resource)
}

func (rsc resource) ListBoards(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	boards, _ := repository.GetAllBoards(conn)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(boards)
}
