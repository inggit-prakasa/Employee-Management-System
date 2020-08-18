package models

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/inggit_prakasa/Employee/database"
)

type Attendance struct {
	Id          int    `json:"id"`
	EmpId       int    `json:"emp_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Created     string `json:"created"`
}

func GetAllAttendance() (Response, error) {
	var obj Attendance
	var arrObj []Attendance
	var res Response

	conn := database.Connection()

	sqlStatement := "SELECT * FROM attendance"

	rows, err := conn.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.EmpId, &obj.Type, &obj.Description, &obj.Created)
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

func AddAttendance(empId int, tipe, description string) (Response, error) {
	var res Response

	v := validator.New()

	att := Attendance{
		EmpId:       empId,
		Type:        tipe,
		Description: description,
	}

	err := v.Struct(att)
	if err != nil {
		return res, err
	}

	conn := database.Connection()

	sqlStatement := "INSERT attendance (attendance_employee_id, attendance_type, attendance_description) VALUES (?,?,?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(empId, tipe, description)
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

func EditAttendance(id, empId int, tipe, description string) (Response, error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "UPDATE attendance SET attendance_employee_id = ?, attendance_type = ?, attendance_description = ? WHERE attendance_id = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(empId, tipe, description, id)
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

func FindAttendance(id int) (Response, error) {
	var res Response
	var obj Attendance

	conn := database.Connection()

	sqlStatement := "SELECT * FROM attendance WHERE attendance_id = ?"

	rows, err := conn.Query(sqlStatement, id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.EmpId, &obj.Type, &obj.Description, &obj.Created)
		if err != nil {
			return res, err
		}

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func DeleteAttendance(id int) (Response, error) {
	var res Response

	conn := database.Connection()

	sqlStatement := "DELETE FROM attendance WHERE attendance_id = ?"

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
