package repository

import (
	"database/sql"

	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/utils"
)

func InsertList(conn *sql.DB, list models.List) (models.List, error) {
	stmt, err := conn.Prepare("insert into list(name, pos, id_board) values (?, ?, ?)")
	utils.ReturnError(err)

	res, err := stmt.Exec(list.Name, list.Order, list.IdBoard)
	utils.ReturnError(err)

	list.Id, err = res.LastInsertId()

	return list, err
}

func GetList(conn *sql.DB, id int) (models.List, error) {
	row := conn.QueryRow("select id, name, pos from list where id = ?", id)
	list := new(models.List)

	err := row.Scan(&list.Id, &list.Name, &list.Order)
	utils.ReturnError(err)

	return *list, err
}

func GetListsByBoard(conn *sql.DB, id int) ([]models.List, error) {
	rows, err := conn.Query("select id, name, pos from list where id_board = ?", id)
	utils.ReturnError(err)
	defer rows.Close()

	var lists []models.List

	for rows.Next() {
		var list models.List
		rows.Scan(&list.Id, &list.Name, &list.Order)
		lists = append(lists, list)
	}
	return lists, err
}

func GetAllLists(conn *sql.DB) ([]models.List, error) {

	rows, err := conn.Query("select id, name, pos from list")
	utils.ReturnError(err)
	defer rows.Close()

	var lists []models.List

	for rows.Next() {
		var list models.List
		rows.Scan(&list.Id, &list.Name, &list.Order)
		lists = append(lists, list)
	}
	return lists, err
}

func UpdateList(conn *sql.DB, list *models.List) (int64, error) {
	res, err := conn.Exec("update list set name = ?, pos = ? where id = ?", list.Name, list.Order, list.Id)
	utils.ReturnError(err)

	rows, err := res.RowsAffected()

	return rows, err
}

func DeleteList(conn *sql.DB, id int) (rows int64, err error) {

	transaction, err := conn.Begin()
	utils.ReturnError(err)

	_, err = transaction.Exec("delete from card where id_list = ?", id)
	utils.ReturnError(err)

	res, err := transaction.Exec("delete from list where id = ?", id)
	utils.ReturnError(err)

	rows, err = res.RowsAffected()
	transaction.Commit()

	return
}
