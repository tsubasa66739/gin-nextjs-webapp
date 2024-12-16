package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/schema"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
)

// ノートを取得する
func GetNote(c *gin.Context) {

	// リクエストバリデーション
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	// 対象のノート取得
	var note repository.TrnNote
	note, err = service.GetNote(uint(id))

	// レスポンスをハンドリングする
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			// 対象のノートが見つからなかった
			c.JSON(http.StatusNotFound, schema.NotFound{
				Message: "Resource not found.",
			})
		} else {
			// 不明なエラーが発生した
			c.JSON(http.StatusInternalServerError, schema.InternalServerError{
				Err:     err,
				Message: "Unknown error.",
			})
		}
		return
	}

	// 正常レスポンス
	c.JSON(http.StatusOK, schema.NoteRes{
		Note:    note,
		Message: "Get note successfully.",
	})
}

func PostNote(c *gin.Context) {

	// リクエストバリデーション
	request := schema.PostNoteReq{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// ノート作成
	note, err := service.CreateNote(&request)

	// レスポンスをハンドリングする
	if err != nil {
		// 不明なエラーが発生した
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// 正常レスポンス
	c.JSON(http.StatusOK, schema.NoteRes{
		Note:    note,
		Message: "Posting note successfully.",
	})
}

func PutNote(c *gin.Context) {

	// リクエストバリデーション
	req := schema.PutNoteReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, schema.BadRequest{
			Message: err.Error(),
		})
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	// ノート更新
	err = service.UpdateNote(uint(id), &req)

	// レスポンスをハンドリングする
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			// 対象のノートが見つからなかった
			c.JSON(http.StatusNotFound, schema.NotFound{
				Message: "Resource not found.",
			})
		} else {
			// 不明なエラーが発生した
			c.JSON(http.StatusInternalServerError, schema.InternalServerError{
				Err:     err,
				Message: "Unknown error.",
			})
		}
		return
	}

	// 正常レスポンス
	c.JSON(http.StatusOK, schema.NoteRes{
		Message: "Putting note successfully.",
	})
}