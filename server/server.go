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

type Handler struct {
	r *mux.Router
}

type Server struct {
	httpServer *http.Server
	Users      db.UserStorage
}

func NewHandler() *Handler {
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
	return &Handler{
		r: r,
	}
}

func (s Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (s Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	infoUser := pkg.User{}

	err = json.Unmarshal(data, &infoUser)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res := FromDomain(s.Users.CreateUser(infoUser))
	inf, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	_, err = w.Write(inf)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s Server) GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, s.Users.U[id])
	w.WriteHeader(http.StatusOK)
}

func (s Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = s.Users.DeleteUser(id)
	w.WriteHeader(http.StatusOK)
}

func (s Server) SubtractBalanceFromUser(w http.ResponseWriter, r *http.Request) {
	infoUser := pkg.User{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &infoUser)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Print(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res := FromDomain(s.Users.UpdateUser(pkg.User{id, infoUser.Name, s.Users.U[id].Balance - infoUser.Balance}))
	inf, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	_, err = w.Write(inf)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s Server) AddBalanceToUser(w http.ResponseWriter, r *http.Request) {
	infoUser := pkg.User{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &infoUser)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	res := FromDomain(s.Users.UpdateUser(pkg.User{id, infoUser.Name, s.Users.U[id].Balance + infoUser.Balance}))
	fmt.Println("result is ---->", res)
	inf, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	_, err = w.Write(inf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	w.WriteHeader(http.StatusCreated)
}
