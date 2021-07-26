package db

import (
	"fmt"
	"log"

	"github.com/sungatuly22/tournament-service/pkg"
)

type UserStorage struct {
	U map[int]pkg.User
}

func (storageUser UserStorage) Tezt(num int) {
	log.Print(num)
}

func (storageUser UserStorage) CreateUser(infoUser pkg.User) pkg.User {
	storageUser.U[infoUser.Id] = pkg.User{Id: infoUser.Id, Name: infoUser.Name, Balance: infoUser.Balance}
	fmt.Println(storageUser.U)
	return storageUser.U[infoUser.Id]
}

func (storageUser UserStorage) UpdateUser(infoUser pkg.User) pkg.User {
	storageUser.U[infoUser.Id] = pkg.User{Id: infoUser.Id, Name: infoUser.Name, Balance: infoUser.Balance}
	return storageUser.U[infoUser.Id]
}

func (storageUser UserStorage) DeleteUser(id int) pkg.User {
	delete(storageUser.U, id)
	return storageUser.U[id]
}
