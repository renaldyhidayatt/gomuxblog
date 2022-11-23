package main

import (
	"log"
	"muxblog/config"
	"muxblog/router"
	"net/http"
)

func main() {

	db, err := config.InitialDatabase()

	if err != nil {
		panic(err.Error())
	}
	router := router.NewCategoryRouter(db)

	log.Fatal(http.ListenAndServe(":5000", router))
}
