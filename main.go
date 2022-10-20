package main

import (
	"crud/config"
	"crud/entities"
	"crud/route"
)

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&entities.News{})

	r := route.SetupRoutes(db)
	r.Run()
}
