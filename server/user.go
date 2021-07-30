package server

import "github.com/sungatuly22/tournament-service/pkg"

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func FromDomain(u pkg.User) User {
	return User{u.Id, u.Name, u.Balance}
}
