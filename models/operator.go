package models

import "fmt"

func GetUserByNicknamePwd(name, pwd string) *User {
	var u User
	rows, err := db.Query("select * from user where name=$1 and password=$2", name, pwd)
	if err != nil{
		fmt.Println("GetUserByNicknamePwd failed", err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next(){
		err = rows.Scan(&u.Id, &u.Nickname, &u.Password, &u.UpdatedAt, &u.Email)
		if err != nil {
			fmt.Println("Get Id failed", err.Error())
			return nil
		}
	}
	fmt.Println(u.Id, u.Nickname, u.Password)
	return &u
}


func GetUserByMailPwd(email, pwd string) *User {
	var u User
	rows, err := db.Query("select * from user where email=$1 and password=$2", email, pwd)
	if err != nil{
		fmt.Println("GetUserByEmailPwd failed", err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next(){
		err = rows.Scan(&u.Id, &u.Nickname, &u.Password, &u.Update, &u.Email)
		if err != nil {
			fmt.Println("Get Id failed", err.Error())
			return nil
		}
	}
	fmt.Println(u.Id, u.Email, u.Password)
	return &u
}

func CreateUserByNamePwdEmail(name, pwd, email string) error{
	_, err := db.Exec("insert into user (name, password, email) values ($1, $2, $3)", name, pwd, email)
	if err != nil{
		fmt.Println("Insert into user failed", err.Error())
		return err
	}
	return nil
}

func CheckUserExistence(name string) bool{
	var u User
	rows, err := db.Query("select * from user where name=$1", name)
	if err != nil{
		fmt.Println("check failed", err.Error())
		return false
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Nickname, &u.Password, &u.UpdatedAt, &u.Email)
		if err != nil {
			fmt.Println("Get UserId failed", err.Error())
			return false
		}
	}

	if u.Id > 0 && u.Nickname != "" && u.Password != ""{
		return true
	}
	return false
}
