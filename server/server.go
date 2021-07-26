package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sungatuly22/tournament-service/db"
	"github.com/sungatuly22/tournament-service/pkg"
)

type Server struct {
	httpServer *http.Server
	Users      db.UserStorage
}

func NewServer() *Server {
	s := &Server{}
	r := mux.NewRouter()
	s.httpServer = &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	s.Users.U = make(map[int]pkg.User)
	r.HandleFunc("/user", s.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", s.GetUserInfoHandler).Methods("GET")
	r.HandleFunc("/user/{id}", s.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/user/{id}/fund", s.AddBalanceToUser).Methods("POST")
	r.HandleFunc("/user/{id}/take", s.SubtractBalanceFromUser).Methods("POST")
	return s
}

func (s Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	infoUser := pkg.User{}

	err = json.Unmarshal(data, &infoUser)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	_ = s.Users.CreateUser(infoUser)

}

func (s Server) GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, s.Users.U[id])
}

func (s Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	_ = s.Users.DeleteUser(id)
}

func (s Server) SubtractBalanceFromUser(w http.ResponseWriter, r *http.Request) {
	infoUser := pkg.User{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	err = json.Unmarshal(data, &infoUser)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	_ = s.Users.UpdateUser(pkg.User{id, infoUser.Name, s.Users.U[id].Balance - infoUser.Balance})
}

func (s Server) AddBalanceToUser(w http.ResponseWriter, r *http.Request) {
	infoUser := pkg.User{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	err = json.Unmarshal(data, &infoUser)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	_ = s.Users.UpdateUser(pkg.User{id, infoUser.Name, s.Users.U[id].Balance + infoUser.Balance})
}