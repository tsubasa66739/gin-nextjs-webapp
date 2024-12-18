package model

type TrnNote struct {
	Model
	Title string
	Body  string
}

type HstNote struct {
	*TrnNote
	NoteID uint
}
