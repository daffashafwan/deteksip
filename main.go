package main

import (
	routes2 "github.com/daffashafwan/deteksip/routes"
)

func main() {

	routes := routes2.Init()

	routes.Logger.Fatal(routes.Start(":9999"))
}
