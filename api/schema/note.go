package schema

import "github.com/tsubasa66739/gin-nextjs-webapp/repository"

type PostNoteReq struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

type PutNoteReq struct {
	Note repository.TrnNote
}

type NoteRes struct {
	Note    repository.TrnNote
	Message string
}
