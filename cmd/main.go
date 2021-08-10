package main

import (
	"log"
	"net/http"

	"github.com/sungatuly22/tournament-service/server"
)

func main() {
	srv := server.NewHandler()
	if err := http.ListenAndServe("localhost:8080", srv); err != nil {
		log.Fatalf(err.Error())
	}
}
