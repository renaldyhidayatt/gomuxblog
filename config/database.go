package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = ""
	database string = "muxblog"
)

var (
	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

func InitialDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}

	return db, err
}
