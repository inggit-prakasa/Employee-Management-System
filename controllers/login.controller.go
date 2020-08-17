package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/inggit_prakasa/Employee/helpers"
	"github.com/inggit_prakasa/Employee/models"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
	"time"
)

type TemplateRenderer struct {
	templates *template.Template
}

type JwtClaims struct {
	Name        string    `json:"name"`
	jwt.StandardClaims
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}


func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username,password)
	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/login")
	}

	cookie := &http.Cookie{}

	cookie.Name = "sessionID"
	cookie.Value = "some_string"
	cookie.Expires = time.Now().Add(48 * time.Hour)

	c.SetCookie(cookie)

	if !res {
		return echo.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	_, err = token.SignedString([]byte("secret"))
	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/login")
	}

	return c.Redirect(http.StatusPermanentRedirect,"/")
}

func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html",nil)
}

func Dashboard(c echo.Context) error {
	result, err := models.GetAllEmployee()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}
	return c.Render(http.StatusOK,"dashboard.html",result)
}


func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}