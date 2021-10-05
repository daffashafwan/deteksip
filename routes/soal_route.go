package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsLoggedInUser = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("user"),
})

var IsLoggedInAnak = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("anak"),
})

func SoalRoute(routes *echo.Echo, api handler.SoalAPI) {

	tps := routes.Group("/soal")
	{
		tps.GET("/list", api.FindAll, IsLoggedInUser)
		tps.GET("/FindByTipeSoalID/:tipesoal", api.FindByTipeSoalID, IsLoggedInAnak, IsLoggedInUser)
		tps.POST("/save", api.SaveOrUpdate, IsLoggedInUser)
	}
}
