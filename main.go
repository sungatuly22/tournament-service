package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	Users.U[Users.lastsID] = User{Id: Users.lastsID, Name: name, Balance: balance}
	Users.lastsID++
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

func MethodCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		CreateUserHandler(w, r)
	case "GET":
		GetUserInfoHandler(w, r)
	case "DELETE":
		DeleteUserHandler(w, r)
	default:
		fmt.Fprint(w, "Please show the correct method")
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var infoUser User
	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		err = json.Unmarshal(data, &infoUser)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		CreateUser(infoUser.Name, infoUser.Balance)
	} else {
		fmt.Fprint(w, "Please, give the correct method")
	}
}

func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'id' is missing")
	}
	key, err := strconv.Atoi(keys[0])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, Users.U[key])
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'id' is missing")
	}
	key, err := strconv.Atoi(keys[0])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	DeleteUser(key)
}

func SubtractBalanceFromUser(w http.ResponseWriter, r *http.Request) {
	var infoUser User
	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		err = json.Unmarshal(data, &infoUser)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		keys, ok := r.URL.Query()["id"]
		if !ok || len(keys[0]) < 1 {
			fmt.Fprintf(w, "Url Param 'id' is missing")
		}
		key, err := strconv.Atoi(keys[0])
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		UpdateUser(key, Users.U[key].Balance-infoUser.Balance)
	} else {
		fmt.Fprint(w, "Please, give the correct method")
	}
}

func AddBalanceToUser(w http.ResponseWriter, r *http.Request) {
	var infoUser User
	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		err = json.Unmarshal(data, &infoUser)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
		keys, ok := r.URL.Query()["id"]
		if !ok || len(keys[0]) < 1 {
			fmt.Fprintf(w, "Url Param 'id' is missing")
		}
		key, err := strconv.Atoi(keys[0])
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		UpdateUser(key, Users.U[key].Balance+infoUser.Balance)
	} else {
		fmt.Fprint(w, "Please, give the correct method")
	}
}

func main() {

	http.HandleFunc("/user", MethodCheck)
	http.HandleFunc("/user/take", SubtractBalanceFromUser)
	http.HandleFunc("/user/fund", AddBalanceToUser)
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}
