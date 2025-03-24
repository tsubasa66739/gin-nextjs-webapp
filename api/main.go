package main

import (
	"log"
	// "net/http"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tsubasa66739/gin-nextjs-webapp/config"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/gorm"
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

	// エンドポイント作成用
	// r := gin.Default()

	// r.POST("/api/item", createItem)

	// r.Run()
}

// type Item struct {
// 	ID          uint   `json:"id"`
// 	Name        string `json:"name"`
// 	Price       int    `json:"price"`
// 	Description string `json:"description"`
// }

// var items []Item

// func createItem(c *gin.Context) {
// 	var newItem Item
// 	if err := c.ShouldBindJSON(&newItem); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	newItem.ID = uint(len(items) + 1)
// 	items = append(items, newItem)
// 	c.JSON(http.StatusCreated, newItem)
// }
