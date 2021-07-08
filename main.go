package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	fmt.Println(Users.U)
}

func UpdateUser(id int, name string, balance int) {
	Users.U[id] = User{Name: name, Balance: balance}
}

func GetUser(id int) {
	log.Println(Users.U[id])
}

func DeleteUser(id int) {
	delete(Users.U, id)
}

func CreateTournament(tournamentName string, deposit int) {

	Tournament.Trs[Tournament.lastsID] = TournamentInfo{id: Tournament.lastsID, name: tournamentName, deposit: deposit}
	Tournament.lastsID++
	log.Println(Tournament)
}

//1 gl obj -> variable empty trcuture []Tournament{}, 0 make
// func UpdateTournaament(TournamentInfo, id) {

// 	if TournamentInfo.name != "" {

// 	}
// }

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var infoUser User
	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(data, &infoUser)
		if err != nil {
			log.Fatal(err)
		}
		CreateUser(infoUser.Name, infoUser.Balance)
	} else {
		fmt.Fprint(w, "Please, give the correct method")
	}
}

// func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		fmt.Println(r.URL.Path)
// 	} else {
// 		fmt.Fprint(w, "Please, give the correct method")
// 	}
// }

func main() {

	http.HandleFunc("/user", CreateUserHandler)
	//http.HandleFunc("/user/id", GetUserInfoHandler)
	//http.HandleFunc("/user/{id}/fund", AddBonusPointsToUser)
	//http.HandleFunc("/user/{id}/take", TakeBonusPointsFromUser)
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}
