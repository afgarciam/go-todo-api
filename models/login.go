package models

type LoginData struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}
