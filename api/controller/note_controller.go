package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller/schema"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
)

type NoteController interface {
	ListNote(c *gin.Context)
	GetNote(c *gin.Context)
	PostNote(c *gin.Context)
	PutNote(c *gin.Context)
}

type noteController struct {
	noteService service.NoteService
}

func NewNoteController(
	noteService service.NoteService,
) NoteController {
	return &noteController{
		noteService: noteService,
	}
}

func (n *noteController) ListNote(c *gin.Context) {
	notes, err := n.noteService.GetNoteList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.InternalServerError{
			Err:     err,
			Message: "Unknown error",
		})
		return
	}

	res := []schema.NoteRes{}
	for _, note := range notes {
		r := schema.NoteRes{
			Title: note.Title,
			Body:  note.Body,
		}
		r.ID = *note.ID
		r.CreatedAt = note.CreatedAt
		r.UpdatedAt = note.UpdatedAt
		res = append(res, r)
	}
	c.JSON(http.StatusOK, res)
}

// ノートを取得する
func (n *noteController) GetNote(c *gin.Context) {

	// リクエストバリデーション
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	// 対象のノート取得
	var note model.TrnNote
	note, err = n.noteService.GetNote(uint(id))

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
	res := schema.NoteRes{
		Title: note.Title,
		Body:  note.Body,
	}
	res.ID = *note.ID
	res.CreatedAt = note.CreatedAt
	res.UpdatedAt = note.UpdatedAt
	c.JSON(http.StatusOK, res)
}

func (n *noteController) PostNote(c *gin.Context) {

	// リクエストバリデーション
	request := schema.PostNoteReq{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// ノート作成
	note, err := n.noteService.CreateNote(&request)

	// レスポンスをハンドリングする
	if err != nil {
		// 不明なエラーが発生した
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// 正常レスポンス
	res := schema.NoteRes{
		Title: note.Title,
		Body:  note.Body,
	}
	res.ID = *note.ID
	res.CreatedAt = note.CreatedAt
	res.UpdatedAt = note.UpdatedAt
	c.JSON(http.StatusOK, res)
}

func (n *noteController) PutNote(c *gin.Context) {

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
	err = n.noteService.UpdateNote(uint(id), &req)

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
	c.JSON(http.StatusOK, gin.H{})
}
