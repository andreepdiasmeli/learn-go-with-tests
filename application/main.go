package main

import (
	"application/server"
	"log"
	"net/http"
)

func main() {
	server := server.NewPlayerServer(server.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
