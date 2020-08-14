package models

import (
	"database/sql"
	"fmt"
	"github.com/inggit_prakasa/Employee/database"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	//var pass string

	con := database.Connection()

	sqlstatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlstatement,username).Scan(
		&obj.Id, &obj.Username,
		)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	return true, nil
}