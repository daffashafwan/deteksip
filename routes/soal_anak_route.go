package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
)


func SoalAnakRoute(routes *echo.Echo, api handler.SoalAPI) {

	tps := routes.Group("/anak/soal")
	{
		tps.GET("/list", api.FindAll, IsLoggedInAnak)
		//tps.GET("/speech", api.InitSpeech)
		tps.POST("/jawab", api.Base64toJpg);
		tps.GET("/FindByID", api.FindByID, IsLoggedInAnak)
		tps.GET("/FindByTipeSoalID", api.FindByTipeSoalID, IsLoggedInAnak)
		tps.POST("/save", api.SaveOrUpdate, IsLoggedInUser)
	}
}
