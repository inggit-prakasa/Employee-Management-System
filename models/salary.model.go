package models

import (
	"github.com/go-playground/validator"
	"github.com/inggit_prakasa/Employee/database"
	"net/http"
)

type Salary struct {
	Id int `json:"id"`
	EmpId int `json:"emp_id"`
	Amount int `json:"amount"`
	Total int `json:"total"`
	Type string `json:"type"`
	Description string `json:"description"`
}

func GetAllSalary() (Response,error) {
	var obj Salary
	var arrObj []Salary
	var res Response

	conn := database.Connection()

	sqlStatement := "SELECT * FROM salary"

	rows, err := conn.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.EmpId, &obj.Amount, &obj.Total, &obj.Type, &obj.Description)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj,obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res,nil
}

func AddSalary(empId,amount,total int,tipe, description string) (Response,error) {
	var res Response

	v := validator.New()

	sal := Salary{
		EmpId:       empId,
		Amount:      amount,
		Total:       total,
		Type:        tipe,
		Description: description,
	}

	err := v.Struct(sal)
	if err != nil {
		return res, err
	}

	conn := database.Connection()

	sqlStatement := "INSERT salary (salary_employee_id, salary_amount, salary_total, salary_type, salary_description) VALUES (?,?,?,?,?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(empId,amount,total,tipe,description)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res,err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64 {
		"last_insert_id" : lastInsertId,
	}

	return res,nil
}

func EditSalary(id, empId, amount,total int, tipe, description string) (Response,error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "UPDATE salary SET salary_employee_id = ?, salary_amount = ?, salary_total = ?, salary_type = ?, salary_description = ? WHERE salary_id = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(empId,amount,total,tipe,description,id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64 {
		"rowsAffected" : rowsAffected,
	}

	return res, nil
}

func DeleteSalary(id int) (Response,error){
	var res Response

	conn := database.Connection()

	sqlStatement := "DELETE FROM salary WHERE salary_id = ?"

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
	res.Data = map[string]int64 {
		"rowsAffected" : rowsAffected,
	}

	return res, nil
}

func FindSalary(id int) (Response, error) {
	var res Response
	var obj Salary

	conn := database.Connection()

	sqlStatement := "SELECT * FROM salary WHERE salary_id = ?"

	rows, err := conn.Query(sqlStatement,id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.EmpId, &obj.Amount, &obj.Total, &obj.Type, &obj.Description)
		if err != nil {
			return res, err
		}

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}