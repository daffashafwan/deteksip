package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
)

func UserRoute(routes *echo.Echo, api handler.UserAPI) {

	mhs := routes.Group("/user")
	{
		mhs.GET("/list", api.FindAll)
		mhs.POST("/save", api.SaveOrUpdate)
		mhs.GET("/FindByUsername/:username", api.FindByUsername)
	}
}
