package schema

import (
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
)

type PostNoteReq struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

type PutNoteReq struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

type NoteRes struct {
	Note    model.TrnNote
	Message string
}
