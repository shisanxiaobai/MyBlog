package model

//管理员帐号
type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
