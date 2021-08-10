package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

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

	res, err := http.NewRequest(http.MethodPost, "/user", r)

	if err != nil {
		t.Fatalf(err.Error())
	}
	recorder := httptest.NewRecorder()

	srv.CreateUserHandler(recorder, res)
	respBody, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	json.Unmarshal(respBody, &result)

	if result.Id != 1 || result.Name != "John" || result.Balance != 950 {
		t.Fatalf("Result is not correct!!!")
	}
}
