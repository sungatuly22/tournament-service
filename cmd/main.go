package main

import (
	"log"

	"github.com/sungatuly22/tournament-service/server"
)

func main() {
	srv := server.NewServer()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}
