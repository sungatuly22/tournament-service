package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TournamentInfo struct {
	id      int
	name    string
	deposit int
	prize   int
	users   []int
	winner  string
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type UserStorage struct {
	U       map[int]User
	lastsID int
}

type TournamentStorage struct {
	Trs     map[int]TournamentInfo
	lastsID int
}

var Users = UserStorage{U: make(map[int]User)}
var Tournament = TournamentStorage{Trs: make(map[int]TournamentInfo)}

func CreateUser(name string, balance int) {
	Users.lastsID++
	Users.U[Users.lastsID] = User{Id: Users.lastsID, Name: name, Balance: balance}
	fmt.Println(Users.U)
}

func UpdateUser(id int, balance int) {
	Users.U[id] = User{Id: id, Name: Users.U[id].Name, Balance: balance}
}

func DeleteUser(id int) {
	delete(Users.U, id)
}

func CreateTournament(tournamentName string, deposit int) {

	Tournament.Trs[Tournament.lastsID] = TournamentInfo{id: Tournament.lastsID, name: tournamentName, deposit: deposit}
	Tournament.lastsID++
	log.Println(Tournament)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var infoUser User
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	err = json.Unmarshal(data, &infoUser)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	CreateUser(infoUser.Name, infoUser.Balance)
}

func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, Users.U[id])
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	DeleteUser(id)
}

func SubtractBalanceFromUser(w http.ResponseWriter, r *http.Request) {
	var infoUser User
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
	UpdateUser(id, Users.U[id].Balance-infoUser.Balance)
}

func AddBalanceToUser(w http.ResponseWriter, r *http.Request) {
	var infoUser User
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
	UpdateUser(id, Users.U[id].Balance+infoUser.Balance)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", GetUserInfoHandler).Methods("GET")
	router.HandleFunc("/user/{id}", DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/user/{id}/fund", AddBalanceToUser).Methods("POST")
	router.HandleFunc("/user/{id}/take", SubtractBalanceFromUser).Methods("POST")
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", router)
}
