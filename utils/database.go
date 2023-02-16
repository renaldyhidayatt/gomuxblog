package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username string = "root"
	password string = ""
	database string = "muxblog"

	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

func Database() (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)

	if err != nil {

		panic(err.Error())
	}

	return conn, nil
}
