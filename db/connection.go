package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ffsales/go-trello-poc/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {

	dbConfig := config.GetDBConfig()
	strConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	fmt.Println(strConn)

	conn, err := sql.Open("mysql", strConn)

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
