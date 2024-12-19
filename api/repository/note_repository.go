package repository

import (
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NoteRepository interface {
	GetById(note *model.TrnNote) error
	Insert(note *model.TrnNote) error
	Update(note *model.TrnNote) error
	InsertHst(note *model.TrnNote) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(
	db *gorm.DB,
) NoteRepository {
	return &noteRepository{
		db: db,
	}
}

func (n *noteRepository) GetById(note *model.TrnNote) error {
	return n.db.Take(&note).Error
}

func (n *noteRepository) Insert(note *model.TrnNote) error {
	// 新規作成して取得する
	return n.db.Clauses(clause.Returning{}).Create(&note).Error
}

func (n *noteRepository) Update(note *model.TrnNote) error {
	// CreatedAt以外を更新する
	return n.db.Select("*").Omit("CreatedAt").Save(&note).Error
}

func (n *noteRepository) InsertHst(note *model.TrnNote) error {
	noteHst := model.HstNote{
		TrnNote: note,
		NoteID:  *note.ID,
	}
	// IDは履歴テーブルの主キーなので除外する
	return n.db.Omit("ID").Create(&noteHst).Error
}
