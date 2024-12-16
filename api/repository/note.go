package repository

import (
	"gorm.io/gorm/clause"
)

type TrnNote struct {
	Model
	Title string
	Body  string
}

type HstNote struct {
	*TrnNote
	NoteID uint
}

func GetNoteById(note *TrnNote) error {
	return db.Take(&note).Error
}

func InsertNote(note *TrnNote) error {
	return db.Clauses(clause.Returning{}).Create(&note).Error
}

func UpdateNote(note *TrnNote) error {
	return db.Select("*").Omit("CreatedAt").Save(&note).Error
}

func InsertNoteHst(note *TrnNote) error {
	noteHst := HstNote{
		TrnNote: note,
		NoteID:  *note.ID,
	}
	return db.Omit("ID").Create(&noteHst).Error
}
