package model

type TrnNote struct {
	Title string
	Body  string
	Model
}

type HstNote struct {
	*TrnNote
	NoteID uint
}
