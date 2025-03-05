package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tsubasa66739/gin-nextjs-webapp/config"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
	"gorm.io/gorm"
	// "github.com/tsubasa66739/gin-nextjs-webapp/service"
	// "github.com/gin-gonic/gin"
)

var db *gorm.DB

func init() {
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
	itemRepository := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepository)
	itemController := controller.NewItemController(itemService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	// r := gin.Default() //１：Defaultルーター　Httpsリクエストを処理し適切なハンドラ関数にルーティング
	// r.GET("/sample", func(c *gin.Context) { //２：エンドポイントのパス、リクエストの処理関数
	// 	c.JSON(200, gin.H{ //ステータスコード、レスポンスボディ
	// 		"message": "pong", //キーと値(map関数)
	// 	})
	// }) //エンドポイントの追加（GETやPOST

	r := gin.Default()
	itemRouter := r.Group("/items")
	authRouter := r.Group("/auth")
	itemRouter.GET("", itemController.FindAll) //関数そのものを渡す（）いらない
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	r.Run("localhost:8080") //３：サーバーのアドレスやポート番号を指定する
}
