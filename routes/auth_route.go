package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	_ "github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/v4/middleware"
	
)

func AuthRoute(routes *echo.Echo, api handler.AuthAPI) {
	mhs := routes.Group("/auth")
	{
		mhs.POST("/login", api.Login)
	}
}
