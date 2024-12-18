package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
	"gorm.io/gorm"
)

// ルーターを初期化する
func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize repositories.
	noteRepository := repository.NewNoteRepository(db)

	// Initialize services.
	noteService := service.NewNoteService(noteRepository)

	// Initialize controllers.
	noteController := NewNoteController(noteService)

	api := r.Group("/api")
	{
		api.GET("/note/:id", noteController.GetNote)
		api.POST("/note", noteController.PostNote)
		api.PUT(("/note/:id"), noteController.PutNote)
	}

	return r
}
