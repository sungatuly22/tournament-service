package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sungatuly22/tournament-service/db"
)

type Server struct {
	httpServer *http.Server
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}

	s.httpServer = &http.Server{
		Addr:    "localhost:8080",
		Handler: s.Router,
	}
	return s
}

// func (s Server) ListenAndServe() error {
// 	return s.httpServer.ListenAndServe()
// }

func (s Server) Routes() {
	s.HandleFunc("/user", db.CreateUserHandler).Methods("POST")
	s.HandleFunc("/user/{id}", db.GetUserInfoHandler).Methods("GET")
	s.HandleFunc("/user/{id}", db.DeleteUserHandler).Methods("DELETE")
	s.HandleFunc("/user/{id}/fund", db.AddBalanceToUser).Methods("POST")
	s.HandleFunc("/user/{id}/take", db.SubtractBalanceFromUser).Methods("POST")
}
