package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
)

func TipeSoalRoute(routes *echo.Echo, api handler.TipeSoalAPI) {

	tps := routes.Group("/tipesoal")
	{
		tps.GET("/list", api.FindAll)
		tps.POST("/save", api.SaveOrUpdate)
	}
}
