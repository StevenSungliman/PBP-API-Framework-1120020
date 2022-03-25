package controller

type User struct {
	Id       int    `form:"id"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Type     int    `form:"type"`
}
