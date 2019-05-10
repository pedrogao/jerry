package model

import (
	password2 "github.com/PedroGao/jerry/libs/password"
	"sync"
)

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[int64]*UserModel
}

type UserModel struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	SayHello string `json:"say_hello"`
	Password string `json:"password"`
}

func (UserModel) IsSuper() bool {
	return false
}

func CreateUser(username, password string) (int64, error) {
	encryptPassword := password2.CreatePassword(password, 5)
	user := &UserModel{
		Username: username,
		Password: string(encryptPassword),
	}
	return DB.InsertOne(user)
}
