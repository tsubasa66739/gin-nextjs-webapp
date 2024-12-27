package schema

type PostNoteReq struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

type PutNoteReq struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

type NoteRes struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	ResSchema
}
