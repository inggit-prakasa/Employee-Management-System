package controllers

import (
	"github.com/inggit_prakasa/Employee/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func AttendancePage(c echo.Context) error {
	return c.Render(http.StatusOK,"attendance.html",nil)
}

func GetAllAttendance(c echo.Context) error {
	result, err := models.GetAllAttendance()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddAttendance(c echo.Context) error {
	empId := c.FormValue("empId")
	tipe := c.FormValue("type")
	description := c.FormValue("description")

	convId, err := strconv.Atoi(empId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.AddAttendance(convId, tipe, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func EditAttendance(c echo.Context) error {
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

	result, err := models.EditAttendance(convId, convEmpId, tipe, desc)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FindAttendance(c echo.Context) error {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.FindAttendance(convId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteAttendance(c echo.Context) error {
	id := c.Param("id")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.DeleteAttendance(convId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
