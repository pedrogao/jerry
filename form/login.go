package form

import (
	"errors"
	"github.com/PedroGao/jerry/libs/password"
	"github.com/PedroGao/jerry/model"
)

// Binding from JSON
type Login struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l Login) ValidateNameAndPassword() error {
	var user = model.UserModel{Username: l.Nickname}
	has, _ := model.DB.Get(&user)
	if !has {
		return errors.New("用户不存在")
	}
	ok := password.ComparePassword([]byte(user.Password), []byte(l.Password))
	if !ok {
		return errors.New("密码错误")
	}
	return nil
}
