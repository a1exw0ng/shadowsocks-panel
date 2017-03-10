package models


func GetUserByNicknamePwd(name, pwd string) *User {
	return &User{Id:1, Nickname:"hello", Password:"password"}
}