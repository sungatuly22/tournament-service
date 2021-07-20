package main

import (
	"net/http"

	"github.com/sungatuly22/tournament-service/server"
)

func main() {
	srv := server.NewServer()
	srv.Routes()
	http.ListenAndServe(":8080", srv)
}
