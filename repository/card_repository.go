package repository

import (
	"database/sql"

	"github.com/ffsales/go-trello-poc/models"
	"github.com/ffsales/go-trello-poc/utils"
)

func InsertCard(conn *sql.DB, card models.Card) (models.Card, error) {

	stmt, err := conn.Prepare("insert into card(name, finished, id_list) values (?, ?, ?)")
	utils.ReturnError(err)
	res, err := stmt.Exec(card.Name, card.Finished, card.IdList)
	utils.ReturnError(err)

	card.Id, err = res.LastInsertId()

	return card, err
}

func GetCard(conn *sql.DB, id int) (models.Card, error) {

	row := conn.QueryRow("select id, name, finished, id_list from card where id = ?", id)
	card := new(models.Card)

	err := row.Scan(&card.Id, &card.Name, &card.Finished, &card.IdList)
	utils.ReturnError(err)

	return *card, err
}

func GetCardsByList(conn *sql.DB, id int) ([]*models.Card, error) {

	rows, err := conn.Query("select id, name, finished, id_list from card where id_list = ?", id)
	utils.ReturnError(err)
	defer rows.Close()

	var cards []*models.Card

	for rows.Next() {
		card := new(models.Card)
		rows.Scan(&card.Id, &card.Name, &card.Finished, &card.IdList)
		cards = append(cards, card)
	}
	return cards, err
}

func GetAllCards(conn *sql.DB) ([]*models.Card, error) {

	rows, err := conn.Query("select id, name, finished, id_list from card")
	utils.ReturnError(err)
	defer rows.Close()

	var cards []*models.Card

	for rows.Next() {
		card := new(models.Card)
		rows.Scan(&card.Id, &card.Name, &card.Finished, &card.IdList)
		cards = append(cards, card)
	}
	return cards, err
}

func UpdateCard(conn *sql.DB, card *models.Card) (int64, error) {
	res, err := conn.Exec("update card set name = ?, finished = ? where id = ?", card.Name, card.Finished, card.Id)
	utils.ReturnError(err)

	rows, err := res.RowsAffected()

	return rows, err
}

func DeleteCard(conn *sql.DB, id int) (rows int64, err error) {
	res, err := conn.Exec("delete from card where id = ?", id)
	utils.ReturnError(err)

	rows, err = res.RowsAffected()

	return
}
