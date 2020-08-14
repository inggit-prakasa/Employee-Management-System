package main

import (
	"github.com/inggit_prakasa/Employee/database"
	"github.com/inggit_prakasa/Employee/routes"
)

func main() {
	database.Connection()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}