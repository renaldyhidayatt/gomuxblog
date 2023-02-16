package main

import (
	"log"
	dbConn "muxblog/db/sqlc"
	"muxblog/router"
	"muxblog/utils"
	"net/http"
)

var (
	db *dbConn.Queries
)

func main() {

	conn, err := utils.Database()

	if err != nil {
		panic(err.Error())
	}

	db = dbConn.New(conn)

	router := router.NewCategoryRouter(db)

	log.Fatal(http.ListenAndServe(":5000", router))
}
