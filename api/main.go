package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tsubasa66739/gin-nextjs-webapp/config"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Setup()
	db = repository.Setup()
}

func main() {
	server := controller.InitRouter(db)
	server.Run()

	items := []model.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", Soldout: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", Soldout: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", Soldout: false},
	}

	itemRepository := repository.NewItemMemoryRepository(items)
	itemService := service.NewItemService(itemRepository)
	itemController := controller.NewItemController(itemService)

	r := gin.Default() //１：Defaultルーター　Httpsリクエストを処理し適切なハンドラ関数にルーティング
	// r.GET("/sample", func(c *gin.Context) { //２：エンドポイントのパス、リクエストの処理関数
	// 	c.JSON(200, gin.H{ //ステータスコード、レスポンスボディ
	// 		"message": "pong", //キーと値(map関数)
	// 	})
	// }) //エンドポイントの追加（GETやPOST

	r.GET("/items", itemController.FindAll) //関数そのものを渡す（）いらない
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run("localhost:8080") //３：サーバーのアドレスやポート番号を指定する

}
