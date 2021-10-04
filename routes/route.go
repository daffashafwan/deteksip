package routes

import (
	db2 "github.com/daffashafwan/deteksip/db"
	"github.com/daffashafwan/deteksip/injection"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	dbConfig := db2.InitDB()

	mahasiswaAPI := injection.InitMahasiswaAPI(dbConfig)
	tipesoalAPI := injection.InitTipeSoalAPI(dbConfig)
	authAPI := injection.InitAuthAPI(dbConfig)
	soalAPI := injection.InitSoalAPI(dbConfig)

	routes := echo.New()

	// set logger
	routes.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, time:${time_unix}, uri=${uri}, status=${status}, error=${error}, latency_human=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out} \n",
	}))

	// Gzip Compression
	routes.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	UserRoute(routes, mahasiswaAPI)
	TipeSoalRoute(routes, tipesoalAPI)
	AuthRoute(routes, authAPI)
	SoalRoute(routes, soalAPI)
	return routes
}
