package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
)



func TipeSoalRoute(routes *echo.Echo, api handler.TipeSoalAPI) {

	tps := routes.Group("/tipesoal")
	{
		tps.GET("/list", api.FindAll, IsLoggedIn)
		tps.POST("/save", api.SaveOrUpdate, IsLoggedIn)
		tps.DELETE("/delete/:id", api.SaveOrUpdate, IsLoggedIn)
	}
}
