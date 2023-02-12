package main

import (
	"database/sql"
	"fmt"
	"log"
	dbConn "muxblog/db/sqlc"
	"muxblog/router"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db       *dbConn.Queries
	username string = "root"
	password string = ""
	database string = "muxblog"

	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

func main() {

	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	db = dbConn.New(conn)

	router := router.NewCategoryRouter(db)

	log.Fatal(http.ListenAndServe(":5000", router))
}
