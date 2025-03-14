package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
	"gorm.io/gorm"
)

// ルーターを初期化する
func InitRouter(db *gorm.DB, items []model.Item) *gin.Engine {
	r := gin.Default()

	// レポジトリの初期化
	noteRepository := repository.NewNoteRepository(db)
	itemRepository := repository.NewItemRepository(db)

	// サービスの初期化
	noteService := service.NewNoteService(noteRepository)
	itemService := service.NewItemService(itemRepository)

	// コントローラの初期化
	noteController := NewNoteController(noteService)
	itemController := NewItemController(itemService)

	//apiのグループ化
	api := r.Group("/api")
	{
		//Note endpoints
		api.GET("/note", noteController.ListNote)
		api.GET("/note/:id", noteController.GetNote)
		api.POST("/note", noteController.PostNote)
		api.PUT(("/note/:id"), noteController.PutNote)

		// Item endpoints
		api.GET("/item", itemController.FindAll)
		api.GET("/item/:id", itemController.FindById)
		api.POST("/item", itemController.Create)
		api.PUT("/item/:id", itemController.Update)
		api.DELETE("/item/:id", itemController.Delete)

	}

	return r
}
