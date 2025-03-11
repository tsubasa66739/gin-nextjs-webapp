package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tsubasa66739/gin-nextjs-webapp/config"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/gorm"
	// "github.com/tsubasa66739/gin-nextjs-webapp/service"
	// "github.com/gin-gonic/gin"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.Setup()
	db = repository.Setup()
}

func main() {
	var items []model.Item

	server := controller.InitRouter(db, items)
	server.Run()

	// items := []model.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", Soldout: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", Soldout: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", Soldout: false},
	// }

	// itemRepository := repository.NewItemMemoryRepository(items)
}
