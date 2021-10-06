package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
)


func SoalRoute(routes *echo.Echo, api handler.SoalAPI) {

	tps := routes.Group("/user/soal")
	{
		tps.GET("/list", api.FindAll, IsLoggedInUser)	
		tps.GET("/FindByID", api.FindByID, IsLoggedInUser)
		tps.GET("/FindByTipeSoalID", api.FindByTipeSoalID, IsLoggedInUser)
		tps.POST("/save", api.SaveOrUpdate, IsLoggedInUser)
		tps.DELETE("/delete", api.DeleteSoal, IsLoggedInUser)
	}
}
