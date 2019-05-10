package form

import "errors"

// Binding from JSON
type Login struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l Login) ValidateNameAndPassword() error {
	if l.Nickname != "pedro" || l.Password != "123456" {
		return errors.New("用户名或密码不正确")
	}
	return nil
}
