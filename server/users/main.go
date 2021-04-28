package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/shubhamjagdhane/myBooksShelf/server/users/api"
)

func main() {

	server, err := api.NewServer()
	if err != nil {
		log.Fatal("Server is not created: ", err)
	}

	err = server.Start()
	if err != nil {
		log.Fatal("Server is not started...")
	}
}
