package main

import (
	"blogspot/config"
	"blogspot/model"
	"blogspot/route"
)

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&model.Post{})


	r:=route.SetupRoute(db)
	r.Run();

}