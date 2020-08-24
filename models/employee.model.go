package models

import (
	"net/http"

	"github.com/inggit_prakasa/Employee/database"

	"github.com/go-playground/validator"
)

type employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Status   string `json:"status"`
	Password string `json:"password"`
}

func GetAllEmployee() (Response, error) {
	var obj employee
	var arrObj []employee
	var res Response

	conn := database.Connection()

	sqlStatement := "SELECT employee_id,employee_name,employee_mobile,employee_email,employee_username,employee_address FROM employee"

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

func AddEmployee(name, mobile, email, username ,password ,address string) (Response, error) {
	var res Response

	v := validator.New()

	emp := employee{
		Name:     name,
		Mobile:   mobile,
		Email:    email,
		Username: username,
		Password: password,
		Address:  address,
	}

	err := v.Struct(emp)
	if err != nil {
		return res, err
	}

	conn := database.Connection()

	sqlStatement := "INSERT employee (employee_name, employee_mobile, employee_email, employee_username,employee_password ,employee_address) VALUES (?,?,?,?,?,?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, mobile, email, username ,password ,address)
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

func CheckEmail(email string) bool {
	conn := database.Connection()
	var jumlah int

	sqlStatement := "SELECT COUNT(employee_id) as 'jumlah' FROM employee WHERE employee_email = ?"

	rows, err := conn.Query(sqlStatement, email)
	defer rows.Close()

	if err != nil {
		return false
	}

	for rows.Next() {
		err = rows.Scan(&jumlah)
		if err != nil {
			return false
		}
	}

	if jumlah == 0 {
		return false
	}
	return true
}

func UpdateEmployee(id int, name, mobile, email, username, address string) (Response, error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "UPDATE employee SET employee_name = ?, employee_mobile = ?, employee_email = ?, employee_username = ?, employee_address = ? WHERE employee_id = ?"

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

func FindEmployee(id int) (Response, error) {
	var res Response
	var obj employee

	conn := database.Connection()

	sqlStatement := "SELECT employee_id,employee_name,employee_mobile,employee_email, employee_username,employee_address FROM employee WHERE employee_id = ?"

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

func DeleteEmployee(id int) (Response, error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "DELETE FROM employee WHERE employee_id = ?"

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
func SetStatusEmployee(id int, status string) (Response, error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "UPDATE employee SET employee_status = ? WHERE employee_id = ?"

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
