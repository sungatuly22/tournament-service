package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sungatuly22/tournament-service/pkg"
)

func TestCreateUserHandler(t *testing.T) {

	srv := &Server{}
	srv.Users.U = make(map[int]pkg.User)

	testUser := User{Id: 1, Name: "John", Balance: 950}
	var result User

	data, err := json.Marshal(testUser)
	if err != nil {
		t.Fatalf(err.Error())
	}

	r := bytes.NewReader(data)

	req, err := http.NewRequest(http.MethodPost, "/user", r)

	if err != nil {
		t.Fatalf(err.Error())
	}
	recorder := httptest.NewRecorder()

	srv.CreateUserHandler(recorder, req)
	respBody, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	json.Unmarshal(respBody, &result)

	if result.Id != 1 || result.Name != "John" || result.Balance != 950 {
		t.Fatalf("Result is not correct!!!")
	}
}

func TestGetUserInfoHandler(t *testing.T) {

	srv := &Server{}
	srv.Users.U = map[int]pkg.User{1: {Id: 1, Name: "John", Balance: 800}, 2: {Id: 2, Name: "Rashford", Balance: 900}}

	result := User{}

	req, err := http.NewRequest(http.MethodGet, "/user/1", nil)

	if err != nil {
		t.Fatalf(err.Error())
	}
	recorder := httptest.NewRecorder()

	srv.GetUserInfoHandler(recorder, req)
	respBody, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	json.Unmarshal(respBody, &result)

	if result != (User{Id: 1, Name: "John", Balance: 800}) {
		t.Fatalf("Result is not correct!!!")
	}
}

func TestDeleteUserHandler(t *testing.T) {
	srv := &Server{}
	srv.Users.U = map[int]pkg.User{1: {Id: 1, Name: "John", Balance: 800}, 2: {Id: 2, Name: "Rashford", Balance: 900}}

	req, err := http.NewRequest(http.MethodDelete, "/user/1", nil)

	if err != nil {
		t.Fatalf(err.Error())
	}
	recorder := httptest.NewRecorder()

	srv.DeleteUserHandler(recorder, req)

	_, ok := srv.Users.U[1]

	if !ok {
		t.Fatalf("Result is not correct!!!")
	}
}

func TestSubtractBalanceFromUser(t *testing.T) {
	srv := &Server{}
	srv.Users.U = map[int]pkg.User{1: {Id: 1, Name: "John", Balance: 800}, 2: {Id: 2, Name: "Rashford", Balance: 900}}

	testUser := User{Id: 1, Name: "John", Balance: 400}
	var result User

	data, err := json.Marshal(testUser)
	if err != nil {
		t.Fatalf(err.Error())
	}

	r := bytes.NewReader(data)

	req, err := http.NewRequest(http.MethodPost, "/user/1/take", r)

	if err != nil {
		t.Fatalf(err.Error())
	}
	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()

	srv.SubtractBalanceFromUser(recorder, req)
	respBody, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	json.Unmarshal(respBody, &result)

	if result != (User{Id: 1, Name: "John", Balance: 400}) {
		fmt.Println(result)
		t.Fatalf("Result is not correct!!!")
	}
}

func TestAddBalanceToUser(t *testing.T) {
	srv := &Server{}
	srv.Users.U = map[int]pkg.User{1: {Id: 1, Name: "John", Balance: 800}, 2: {Id: 2, Name: "Rashford", Balance: 900}}

	testUser := pkg.User{Id: 1, Name: "John", Balance: 400}
	var result pkg.User

	data, err := json.Marshal(testUser)
	if err != nil {
		t.Fatalf(err.Error())
	}

	r := bytes.NewReader(data)

	req, err := http.NewRequest(http.MethodPost, "/user/1/fund", r)

	if err != nil {
		t.Fatalf(err.Error())
	}
	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()

	srv.AddBalanceToUser(recorder, req)
	respBody, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	json.Unmarshal(respBody, &result)

	if result != (pkg.User{Id: 1, Name: "John", Balance: 1200}) {
		fmt.Println(result)
		t.Fatalf("Result is not correct!!!")
	}
}
