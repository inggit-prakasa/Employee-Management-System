package controllers

import (
	"net/http"
	"strconv"

	"github.com/inggit_prakasa/Employee/models"
	"github.com/labstack/echo"
)

func SalaryPage(c echo.Context) error {
	return c.Render(http.StatusOK, "salary.html", nil)
}

func GetAllSalary(c echo.Context) error {
	result, err := models.GetAllSalary()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddSalary(c echo.Context) error {
	employee_id := c.FormValue("employee_id")
	amount := c.FormValue("amount")
	total := c.FormValue("total")
	tipe := c.FormValue("type")
	description := c.FormValue("description")

	convEmpId, err := strconv.Atoi(employee_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	convAmount, err := strconv.Atoi(amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	convTotal, err := strconv.Atoi(total)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.AddSalary(convEmpId, convAmount, convTotal, tipe, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func EditSalary(c echo.Context) error {
	id := c.FormValue("id")
	employee_id := c.FormValue("employee_id")
	amount := c.FormValue("amount")
	total := c.FormValue("total")
	tipe := c.FormValue("type")
	description := c.FormValue("description")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	convEmpId, err := strconv.Atoi(employee_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	convAmount, err := strconv.Atoi(amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	convTotal, err := strconv.Atoi(total)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.EditSalary(convId, convEmpId, convAmount, convTotal, tipe, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSalary(c echo.Context) error {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.DeleteSalary(convId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindSalary(c echo.Context) error {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.FindSalary(convId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
