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

func (note *TrnNote) GetById() error {
	return db.Take(&note).Error
}

func (note *TrnNote) Insert() error {
	// 新規作成して取得する
	return db.Clauses(clause.Returning{}).Create(&note).Error
}

func (note *TrnNote) Update() error {
	// CreatedAt以外を更新する
	return db.Select("*").Omit("CreatedAt").Save(&note).Error
}

func (note *TrnNote) InsertHst() error {
	noteHst := HstNote{
		TrnNote: note,
		NoteID:  *note.ID,
	}
	// IDは履歴テーブルの主キーなので除外する
	return db.Omit("ID").Create(&noteHst).Error
}
