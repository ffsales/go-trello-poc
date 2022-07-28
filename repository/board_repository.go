package repository

import (
	"database/sql"

	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/utils"
)

func InsertBoard(conn *sql.DB, board models.Board) (models.Board, error) {
	stmt, err := conn.Prepare("insert into board(name, description) values (?, ?)")
	utils.ReturnError(err)

	res, err := stmt.Exec(board.Name, board.Description)
	utils.ReturnError(err)

	board.Id, err = res.LastInsertId()

	return board, err
}

func GetBoard(conn *sql.DB, id int) (models.Board, error) {
	row := conn.QueryRow("select id, name, description from board where id = ?", id)
	board := new(models.Board)

	err := row.Scan(&board.Id, &board.Name, &board.Description)
	utils.ReturnError(err)

	return *board, err
}

func GetAllBoards(conn *sql.DB) ([]models.Board, error) {

	rows, err := conn.Query("select id, name, description from board")
	utils.ReturnError(err)
	defer rows.Close()

	var boards []models.Board

	for rows.Next() {
		var board models.Board
		rows.Scan(&board.Id, &board.Name, &board.Description)
		boards = append(boards, board)
	}
	return boards, err
}

func UpdateBoard(conn *sql.DB, board *models.Board) (int64, error) {
	res, err := conn.Exec("update board set name = ?, description = ? where id = ?", board.Name, board.Description, board.Id)
	utils.ReturnError(err)

	rows, err := res.RowsAffected()

	return rows, err
}

func DeleteBoard(conn *sql.DB, id int) (rows int64, err error) {
	transaction, err := conn.Begin()
	utils.ReturnError(err)

	_, err = transaction.Exec("delete c from card c inner join list l on c.id_list = l.id where l.id_board = ?", id)
	utils.ReturnError(err)

	_, err = transaction.Exec("delete from list where id_board = ?", id)
	utils.ReturnError(err)

	res, err := transaction.Exec("delete from board where id = ?", id)
	utils.ReturnError(err)

	rows, err = res.RowsAffected()

	transaction.Commit()

	return
}
