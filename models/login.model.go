package models

import (
	"database/sql"
	"fmt"

	"github.com/inggit_prakasa/Employee/database"
	"github.com/inggit_prakasa/Employee/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(email, password string) (bool, error) {
	var obj User
	var pass string

	con := database.Connection()

	sqlstatement := "SELECT employee_id,employee_username,employee_password FROM employee WHERE employee_email = ?"

	err := con.QueryRow(sqlstatement, email).Scan(
		&obj.Id, &obj.Username, &pass,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pass)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
