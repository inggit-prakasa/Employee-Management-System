package controllers

import (
	"net/http"
	"strconv"

	"github.com/inggit_prakasa/Employee/models"

	"github.com/labstack/echo"
)

func LeavePage(c echo.Context) error {
	return c.Render(http.StatusOK, "leave.html", nil)
}

func GetAllLeave(c echo.Context) error {
	result, err := models.GetAllLeave()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddLeave(c echo.Context) error {
	empId := c.FormValue("empId")
	tipe := c.FormValue("type")
	description := c.FormValue("description")

	convId, err := strconv.Atoi(empId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.AddLeave(convId, tipe, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func EditLeave(c echo.Context) error {
	id := c.FormValue("id")
	empId := c.FormValue("empId")
	tipe := c.FormValue("type")
	desc := c.FormValue("description")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	convEmpId, err := strconv.Atoi(empId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.EditLeave(convId, convEmpId, tipe, desc)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindLeave(c echo.Context) error {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.FindLeave(convId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteLeave(c echo.Context) error {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.DeleteLeave(convId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
