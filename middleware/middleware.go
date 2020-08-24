package middleware

import (
	"github.com/labstack/echo/middleware"
)

var (
	CookieSessionLogin    = "SessionLogin"
	LoginSuccessCookieVal = "LoginSuccess"
	UserIdExample         = "20022012"
	SecretKeyExample      = "mLmHu8f1IxFo4dWurBG3jEf1Ex0wDZvvwND6eFmcaX"
	SigningMethodExample  = "HS512"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

var TokenJwt = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: SigningMethodExample,
	SigningKey:    []byte(SecretKeyExample),
})