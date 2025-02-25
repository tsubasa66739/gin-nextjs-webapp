package main

import (
	"github.com/tsubasa66739/gin-nextjs-webapp/config"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Setup()
	db = repository.Setup()
}

func main() {
	// items := []model.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	// }

	//server := controller.InitRouter(db)
	var items []model.Item

	server := controller.InitRouter(db, items)
	server.Run()
}
