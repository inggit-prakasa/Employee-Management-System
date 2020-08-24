package controllers

import (
	"fmt"
	"github.com/inggit_prakasa/Employee/helpers"
	"net/http"
	"strconv"

	"github.com/inggit_prakasa/Employee/models"
	"github.com/labstack/echo"
)

func EmployeePage(c echo.Context) error {
	result, err := models.GetAllEmployee()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.Render(http.StatusOK, "employee.html", result)
}

func GetAllEmployee(c echo.Context) error {
	result, err := models.GetAllEmployee()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddEmployee(c echo.Context) error {
	name := c.FormValue("name")
	mobile := c.FormValue("mobile")
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")
	address := c.FormValue("address")

	hashPass, _ := helpers.HashPassword(password)

	if models.CheckEmail(email) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Email sudah ada"})
	}

	result, err := models.AddEmployee(name, mobile, email, username, hashPass ,address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddEmployeePage(c echo.Context) error {
	return c.Render(http.StatusOK,"addemployee.html",nil)
}

func UpdateEmployee(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	mobile := c.FormValue("mobile")
	email := c.FormValue("email")
	username := c.FormValue("username")
	address := c.FormValue("address")
	fmt.Println(id)
	fmt.Println(name)
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.UpdateEmployee(convId, name, mobile, email, username, address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteEmployee(c echo.Context) error {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.DeleteEmployee(convId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindEmployee(c echo.Context) error {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.FindEmployee(convId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Render(http.StatusOK, "editemployee.html",result)
}
