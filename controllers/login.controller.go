package controllers

import (
	"github.com/inggit_prakasa/Employee/helpers"
	"github.com/inggit_prakasa/Employee/models"
	"github.com/labstack/echo"
	"net/http"
)

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username,password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	GenerateHashPassword(c)

	return c.JSON(http.StatusOK, map[string]string{

		"password": password,
	})
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}