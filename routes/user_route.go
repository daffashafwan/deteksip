package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
)

func UserRoute(routes *echo.Echo, api handler.UserAPI) {

	mhs := routes.Group("/user")
	{
		mhs.GET("/list", api.FindAll, IsLoggedIn)
		mhs.POST("/save", api.SaveOrUpdate, IsLoggedIn)
		mhs.GET("/FindByUsername/:username", api.FindByUsername, IsLoggedIn)
	}
}
