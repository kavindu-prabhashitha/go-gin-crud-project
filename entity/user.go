package entity

type LoginUser struct {
	UserName string `json:"username" binding:"required"`
	Passowrd string `json:"password" binding:"required"`
}
