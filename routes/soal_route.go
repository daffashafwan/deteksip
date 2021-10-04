package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
)


func SoalRoute(routes *echo.Echo, api handler.SoalAPI) {

	tps := routes.Group("/soal")
	{
		tps.GET("/list", api.FindAll)
		tps.GET("/FindByTipeSoalID/:tipesoal", api.FindByTipeSoalID)
		tps.POST("/save", api.SaveOrUpdate)
	}
}
