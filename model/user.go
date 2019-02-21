package model

import (
	"sync"
)

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserModel
}

type UserModel struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	SayHello string `json:"say_hello"`
	Password string `json:"password"`
}

func (UserModel) IsSuper() bool {
	return false
}
