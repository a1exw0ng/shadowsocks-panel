package models

import (
	"database/sql"
	"sync"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var m *sync.Mutex
func InitDB(path string) {
	var err error
	db, err = sql.Open("sqlite3", path)
	//defer db.Close()

	if err != nil{
		//return nil, errors.New("Open database failed")
		fmt.Println("Open database failed", err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database open success")

	m = new(sync.Mutex)

}