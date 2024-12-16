package controller

import "github.com/gin-gonic/gin"

// ルーターを初期化する
func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/note/:id", GetNote)
		api.POST("/note", PostNote)
		api.PUT(("/note/:id"), PutNote)
	}

	return r
}
