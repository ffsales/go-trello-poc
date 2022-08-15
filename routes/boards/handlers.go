package boards

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ffsales/go-trello-poc/db"
	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/repository"
	"github.com/ffsales/go-trello-poc/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetAllBoards(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	boards, err := repository.GetAllBoards(conn)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Error Service")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	respBoards := []render.Renderer{}

	for _, board := range boards {
		respBoards = append(respBoards, board.ToResponse())
	}

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, respBoards)
}

func GetBoard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strBoardId := chi.URLParam(r, "boardId")
	boardId, err := strconv.Atoi(strBoardId)
	if err != nil {
		utils.BadRequestError(w, r, err, "Invalid Id")
		return
	}

	board, err := repository.GetBoard(conn, boardId)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern Error")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, board.ToResponse())
}

func CreateBoard(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Body == nil {
		utils.UnprocessableEntityError(w, r, err, "Empty body")
		return
	}

	conn := db.GetConnection()
	defer conn.Close()

	var requestBoard models.Board
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBoard); err != nil {
		utils.UnprocessableEntityError(w, r, err, "Body incorrect")
		return
	}

	board, err := repository.InsertBoard(conn, requestBoard)
	if err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern Error")
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, board.ToResponse())
}

func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Body == nil {
		utils.BadRequestError(w, r, err, "Body incorrect")
		return
	}

	conn := db.GetConnection()
	defer conn.Close()

	strBoardId := chi.URLParam(r, "boardId")
	boardId, err := strconv.Atoi(strBoardId)
	if err != nil {
		utils.BadRequestError(w, r, err, "Invalid Id")
		return
	}

	var requestBoard models.Board
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBoard); err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern Error")
		return
	}

	foundBoard, err := repository.GetBoard(conn, int(boardId))
	if err != nil {
		utils.NotFoundError(w, r, err, "Intern Error")
		return
	}

	foundBoard.Name = requestBoard.Name
	foundBoard.Description = requestBoard.Description

	if rows, err := repository.UpdateBoard(conn, &foundBoard); err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern Error")
		return
	} else if rows != 1 {
		utils.BadRequestError(w, r, err, fmt.Sprintf("Error: %d rows affected", rows))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, foundBoard.ToResponse())
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	conn := db.GetConnection()
	defer conn.Close()

	strBoardId := chi.URLParam(r, "boardId")
	boardId, err := strconv.Atoi(strBoardId)
	if err != nil {
		utils.NotFoundError(w, r, err, "Intern Error")
		return
	}

	if rows, err := repository.DeleteBoard(conn, boardId); err != nil {
		utils.ServiceUnavailableError(w, r, err, "Intern Error")
		return
	} else if rows != 1 {
		log.Printf("Error: %d rows affected", rows)
	}

	w.WriteHeader(http.StatusNoContent)
}
