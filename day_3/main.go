package main

import (
	"day_2/config"
	"day_2/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()
	e := routes.New()

	//eAuth := e.Group("")
	//eAuth.Use(echoMid.BasicAuth(middlewares.BasicAuthDb))
	e.Logger.Fatal(e.Start(":9000"))

}
