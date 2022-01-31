package Conn

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB

func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", "root:@/member?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)

	}

	return db

}
