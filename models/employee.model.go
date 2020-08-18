package models

import (
	"github.com/inggit_prakasa/Employee/database"
	"net/http"

	"github.com/go-playground/validator"
)

type github.com/inggit_prakasa/Employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Status   string `json:"status"`
}

func GetAllgithub.com/inggit_prakasa/Employee() (Response, error) {
	var obj github.com/inggit_prakasa/Employee
	var arrObj []github.com/inggit_prakasa/Employee
	var res Response

	conn := database.Connection()

	sqlStatement := "SELECT github.com/inggit_prakasa/Employee_id,github.com/inggit_prakasa/Employee_name,github.com/inggit_prakasa/Employee_mobile,github.com/inggit_prakasa/Employee_email, github.com/inggit_prakasa/Employee_username,github.com/inggit_prakasa/Employee_address FROM github.com/inggit_prakasa/Employee"

	rows, err := conn.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Mobile, &obj.Email, &obj.Username, &obj.Address)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func Addgithub.com/inggit_prakasa/Employee(name, mobile, email, username, address string) (Response, error) {
	var res Response

	v := validator.New()

	emp := github.com/inggit_prakasa/Employee{
		Name:     name,
		Mobile:   mobile,
		Email:    email,
		Username: username,
		Address:  address,
	}

	err := v.Struct(emp)
	if err != nil {
		return res, err
	}

	conn := database.Connection()

	sqlStatement := "INSERT github.com/inggit_prakasa/Employee (github.com/inggit_prakasa/Employee_name, github.com/inggit_prakasa/Employee_mobile, github.com/inggit_prakasa/Employee_email, github.com/inggit_prakasa/Employee_username, github.com/inggit_prakasa/Employee_address) VALUES (?,?,?,?,?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, mobile, email, username, address)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_insert_id": lastInsertId,
	}

	return res, nil

}

func Updategithub.com/inggit_prakasa/Employee(id int, name, mobile, email, username, address string) (Response, error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "UPDATE github.com/inggit_prakasa/Employee SET github.com/inggit_prakasa/Employee_name = ?, github.com/inggit_prakasa/Employee_mobile = ?, github.com/inggit_prakasa/Employee_email = ?, github.com/inggit_prakasa/Employee_username = ?, github.com/inggit_prakasa/Employee_address = ? WHERE github.com/inggit_prakasa/Employee_id = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, mobile, email, username, address, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

func Findgithub.com/inggit_prakasa/Employee(id int) (Response, error) {
	var res Response
	var obj github.com/inggit_prakasa/Employee

	conn := database.Connection()

	sqlStatement := "SELECT github.com/inggit_prakasa/Employee_id,github.com/inggit_prakasa/Employee_name,github.com/inggit_prakasa/Employee_mobile,github.com/inggit_prakasa/Employee_email, github.com/inggit_prakasa/Employee_username,github.com/inggit_prakasa/Employee_address FROM github.com/inggit_prakasa/Employee WHERE github.com/inggit_prakasa/Employee_id = ?"

	rows, err := conn.Query(sqlStatement, id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Mobile, &obj.Email, &obj.Username, &obj.Address)
		if err != nil {
			return res, err
		}

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func Deletegithub.com/inggit_prakasa/Employee(id int) (Response, error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "DELETE FROM github.com/inggit_prakasa/Employee WHERE github.com/inggit_prakasa/Employee_id = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

//------------------------------------------------------------------------------------------
func SetStatusgithub.com/inggit_prakasa/Employee(id int, status string) (Response, error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "UPDATE github.com/inggit_prakasa/Employee SET github.com/inggit_prakasa/Employee_status = ? WHERE github.com/inggit_prakasa/Employee_id = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(status, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}
