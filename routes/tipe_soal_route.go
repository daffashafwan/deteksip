package routes

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("admin"),
})

func TipeSoalRoute(routes *echo.Echo, api handler.TipeSoalAPI) {

	tps := routes.Group("/tipesoal")
	{
		tps.GET("/list", api.FindAll, IsLoggedIn)
		tps.POST("/save", api.SaveOrUpdate, IsLoggedIn)
	}
}
