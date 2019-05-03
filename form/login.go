package form

import "errors"

// Binding from JSON
type Login struct {
	NickName string `form:"nickname" json:"nickname" xml:"nickname"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func (l Login) ValidateNameAndPassword() error {
	if l.NickName != "pedro" || l.Password != "123456" {
		return errors.New("用户名或密码不正确")
	}
	return nil
}
