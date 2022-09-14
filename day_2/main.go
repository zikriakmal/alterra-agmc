package main

import (
	"day_2/config"
	"day_2/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()
	e.Logger.Fatal(e.Start(":9000"))
}
