package repository

import (
	"database/sql"

	"github.com/ffsales/go-trello-poc/models"
)

func InsertBoard(conn *sql.DB, board models.Board) (models.Board, error) {
	stmt, err := conn.Prepare("insert into board(name, description) values (?, ?)")
	if err != nil {
		return board, err
	}

	res, err := stmt.Exec(board.Name, board.Description)
	if err != nil {
		return board, err
	}

	board.Id, err = res.LastInsertId()

	return board, err
}

func GetBoard(conn *sql.DB, id int) (models.Board, error) {
	row := conn.QueryRow("select id, name, description from board where id = ?", id)
	board := new(models.Board)

	err := row.Scan(&board.Id, &board.Name, &board.Description)
	if err != nil {
		return models.Board{}, err
	}

	return *board, err
}

func GetAllBoards(conn *sql.DB) ([]*models.Board, error) {
	var boards []*models.Board
	rows, err := conn.Query("select id, name, description from board")
	if err != nil {
		return boards, err
	}
	defer rows.Close()

	for rows.Next() {
		board := new(models.Board)
		rows.Scan(&board.Id, &board.Name, &board.Description)
		boards = append(boards, board)
	}
	return boards, err
}

func UpdateBoard(conn *sql.DB, board *models.Board) (int64, error) {
	res, err := conn.Exec("update board set name = ?, description = ? where id = ?", board.Name, board.Description, board.Id)
	if err != nil {
		return 0, err
	}
	rows, err := res.RowsAffected()

	return rows, err
}

func DeleteBoard(conn *sql.DB, id int) (rows int64, err error) {
	transaction, err := conn.Begin()
	if err != nil {
		return 0, err
	}

	_, err = transaction.Exec("delete c from card c inner join list l on c.id_list = l.id where l.id_board = ?", id)
	if err != nil {
		return 0, err
	}

	_, err = transaction.Exec("delete from list where id_board = ?", id)
	if err != nil {
		return 0, err
	}

	res, err := transaction.Exec("delete from board where id = ?", id)
	if err != nil {
		return 0, err
	}

	rows, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	transaction.Commit()

	return
}
