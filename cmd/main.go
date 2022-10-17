package main

import (
	"main/internal/api"
	"main/internal/database"
)

func main() {
	database.InitDB()
	api.InitServer()
}
