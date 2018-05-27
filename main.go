package main

import (
	"github.com/bene/dockyard/server"
	"log"
	"net/http"
)

func main() {

	router := server.NewServer()

	err := http.ListenAndServe(":5050", router)
	if err != nil {
		log.Fatalln(err)
	}
}
