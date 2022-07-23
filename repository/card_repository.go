package repository

import (
	"database/sql"
	"log"

	"github.com/ffsales/go-trello-poc/models"
)

func Insert(conn *sql.DB, card models.Card) (models.Card, error) {

	stmt, err := conn.Prepare("insert into card(name, finished) values (?, ?)")
	returnError(err)
	res, err := stmt.Exec(card.Name, card.Finished)
	returnError(err)

	card.Id, err = res.LastInsertId()

	return card, err
}

func Get(conn *sql.DB, id int) (models.Card, error) {

	row := conn.QueryRow("select id, name, finished from card where id = ?", id)
	card := new(models.Card)

	err := row.Scan(&card.Id, &card.Name, &card.Finished)
	returnError(err)

	return *card, err
}

func GetAll(conn *sql.DB) ([]models.Card, error) {

	rows, err := conn.Query("select id, name, finished from card")
	returnError(err)
	defer rows.Close()

	var cards []models.Card

	for rows.Next() {
		var card models.Card
		rows.Scan(&card.Id, &card.Name, &card.Finished)
		cards = append(cards, card)
	}
	return cards, err
}

func Update(conn *sql.DB, card *models.Card) (int64, error) {
	res, err := conn.Exec("update card set name = ?, finished = ? where id = ?", card.Name, card.Finished, card.Id)
	returnError(err)

	rows, err := res.RowsAffected()

	return rows, err
}

func Delete(conn *sql.DB, id int) (rows int64, err error) {
	res, err := conn.Exec("delete from card where id = ?", id)
	returnError(err)

	rows, err = res.RowsAffected()

	return
}

func returnError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
