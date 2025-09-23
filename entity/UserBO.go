package entity

type UserBO struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
