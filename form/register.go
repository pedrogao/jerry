package form

type Register struct {
	Nickname        string `json:"nickname" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=Password"`
}
