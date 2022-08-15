package repository

import (
	"database/sql"

	"github.com/ffsales/go-trello-poc/models"
)

func InsertList(conn *sql.DB, list models.List) (models.List, error) {
	stmt, err := conn.Prepare("insert into list(name, pos, id_board) values (?, ?, ?)")
	if err != nil {
		return list, err
	}

	res, err := stmt.Exec(list.Name, list.Order, list.IdBoard)
	if err != nil {
		return list, err
	}

	list.Id, err = res.LastInsertId()

	return list, err
}

func GetList(conn *sql.DB, id int) (models.List, error) {
	row := conn.QueryRow("select id, name, pos, id_board from list where id = ?", id)
	list := new(models.List)

	err := row.Scan(&list.Id, &list.Name, &list.Order, &list.IdBoard)
	if err != nil {
		return *list, err
	}

	return *list, err
}

func GetListsByBoard(conn *sql.DB, id int) ([]*models.List, error) {
	var lists []*models.List
	rows, err := conn.Query("select id, name, pos, id_board from list where id_board = ?", id)
	if err != nil {
		return lists, err
	}
	defer rows.Close()

	for rows.Next() {
		list := new(models.List)
		rows.Scan(&list.Id, &list.Name, &list.Order, &list.IdBoard)
		lists = append(lists, list)
	}
	return lists, err
}

func GetAllLists(conn *sql.DB) ([]*models.List, error) {
	var lists []*models.List

	rows, err := conn.Query("select id, name, pos, id_board from list")
	if err != nil {
		return lists, err
	}
	defer rows.Close()

	for rows.Next() {
		list := new(models.List)
		rows.Scan(&list.Id, &list.Name, &list.Order, &list.IdBoard)
		lists = append(lists, list)
	}
	return lists, err
}

func UpdateList(conn *sql.DB, list *models.List) (int64, error) {
	res, err := conn.Exec("update list set name = ?, pos = ? where id = ?", list.Name, list.Order, list.Id)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()

	return rows, err
}

func DeleteList(conn *sql.DB, id int) (rows int64, err error) {

	transaction, err := conn.Begin()
	if err != nil {
		return 0, err
	}

	_, err = transaction.Exec("delete from card where id_list = ?", id)
	if err != nil {
		return 0, err
	}

	res, err := transaction.Exec("delete from list where id = ?", id)
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
