package db

import (
	"testing"

	"github.com/sungatuly22/tournament-service/pkg"
)

func TestCreateUser(t *testing.T) {
	inf := UserStorage{make(map[int]pkg.User)}
	tests := []struct {
		name     string
		id       int
		username string
		balance  int
	}{
		{
			name:     "#1 success",
			id:       1,
			username: "John",
			balance:  400,
		},
		{
			name:     "#2 success",
			id:       2,
			username: "Barbosa",
			balance:  550,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if infoUser := inf.CreateUser(pkg.User{test.id, test.username, test.balance}); infoUser != (pkg.User{test.id, test.name, test.balance}) {
				t.Fatalf("Expected username or expected balance is not correst!!!!")
			}
			_ = inf.DeleteUser(test.id)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	inf := UserStorage{make(map[int]pkg.User)}
	tests := []struct {
		name string
		id   int
	}{
		{
			name: "#1 success",
			id:   1,
		},
		{
			name: "#2 success",
			id:   2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if infoUser := inf.DeleteUser(test.id); infoUser != (pkg.User{}) {
				t.Fatalf("User is not deleted!!!!")
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	inf := UserStorage{make(map[int]pkg.User)}
	tests := []struct {
		name     string
		id       int
		username string
		balance  int
	}{
		{
			name:     "#1 success",
			id:       1,
			username: "John",
			balance:  400,
		},
		{
			name:     "#2 success",
			id:       3,
			username: "Barbosa",
			balance:  550,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if infoUser := inf.UpdateUser(pkg.User{test.id, test.username, test.balance}); infoUser != (pkg.User{test.id, test.username, test.balance}) {
				t.Fatalf("Expected username or expected balance is not correst!!!!")
			}
		})
	}
}
