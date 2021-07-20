package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type UserStorage struct {
	U       map[int]User
	lastsID int
}

var Users = UserStorage{U: make(map[int]User)}

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
