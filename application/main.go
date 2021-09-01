package main

import (
	"application/server"
	"log"
	"net/http"
)

func main() {
	server := &server.PlayerServer{Store: server.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
