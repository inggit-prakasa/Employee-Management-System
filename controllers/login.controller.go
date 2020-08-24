package controllers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/inggit_prakasa/Employee/helpers"
	"github.com/inggit_prakasa/Employee/models"
	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var (
	CookieSessionLogin    = "SessionLogin"
	LoginSuccessCookieVal = "LoginSuccess"
	UserIdExample         = "20022012"
	SecretKeyExample      = "mLmHu8f1IxFo4dWurBG3jEf1Ex0wDZvvwND6eFmcaX"
	SigningMethodExample  = "HS512"
)

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func CheckLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := models.CheckLogin(email, password)
	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/login")
	}

	if !res {
		return echo.ErrUnauthorized
	}

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, "/login")
	}


	cookie := &http.Cookie{}

	cookie.Name = CookieSessionLogin
	cookie.Value = t
	cookie.Expires = time.Now().Add(48 * time.Hour)

	c.SetCookie(cookie)

	return c.Redirect(http.StatusPermanentRedirect, "/admin/dashboard")
}

func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func Dashboard(c echo.Context) error {
	result, err := models.GetAllEmployee()
	cookie, err := c.Cookie(CookieSessionLogin)
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			return c.String(http.StatusUnauthorized, "WARNING: You don't have any cookie")
		}
		log.Println(err)
		return err
	}
	tknStr := cookie.Value

	claims := &Claims{}

	_, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	//result["Data"] = claims.Email

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.Render(http.StatusOK, "dashboard.html", result)
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func LogOut(c echo.Context) error {
	cookie := &http.Cookie{}

	cookie.Name = CookieSessionLogin
	cookie.Value = ""
	cookie.MaxAge = -1

	c.SetCookie(cookie)
	fmt.Println(cookie)
	return c.Redirect(http.StatusPermanentRedirect, "/login")
}

func CheckCookieLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(CookieSessionLogin)
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "WARNING: You don't have any cookie")
			}
			log.Println(err)
			return err
		}
		tknStr := cookie.Value

		claims := &Claims{}

		_, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims.Email != "" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "WARNING: You don't have the right cookie")
	}
}



func MainJwt(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		claims := token.Claims.(jwt.MapClaims)
		log.Println("User Name: ", claims["name"], ", User ID: ", claims["jti"])
		return c.JSON(http.StatusOK, map[string] interface {} {
		"claims": claims,
		"message": "SUCCESS: you are on the top secret jwt page!",
	})
}

func CreateJWT(c echo.Context) error {
	token, err := CreateJwtToken()
	if err != nil {
		log.Println("Error when creating JWT token", err)
		return c.String(http.StatusInternalServerError, "ERROR: something went wrong while creating JWT token!")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "You ware logged in!",
		"token":   token,
	})
}

func CreateJwtToken() (string, error) {
	claims := JwtClaims{
		"Firman",
		jwt.StandardClaims{
			Id:        UserIdExample,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	// we hash the jwt claims
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(SecretKeyExample))
	if err != nil {
		return "", err
	}

	return token, nil
}

