package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/v4/middleware"
)

func AuthRoute(routes *echo.Echo, api handler.AuthAPI) {

	mhs := routes.Group("/auth")
	{
		mhs.POST("/login", api.Login)
	}
	config := middleware.JWTConfig{
		Claims:     &jwt,
		SigningKey: []byte("secret"),
	}
	mhs.Use(middleware.JWTWithConfig(config))

}
