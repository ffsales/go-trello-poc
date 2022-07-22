package repository

import (
	"database/sql"
	"fmt"

	"github.com/ffsales/go-trello-poc/models"
)

func Insert(conn *sql.DB, card models.Card) error {

	stmt, errPrepare := conn.Prepare("insert into card(name, finished) values (?, ?)")
	returnError(errPrepare)
	res, err := stmt.Exec(card.Name, card.Finished)
	returnError(err)

	card.Id, err = res.LastInsertId()

	fmt.Println(card)

	return err
}

func returnError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
