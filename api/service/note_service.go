package service

import (
	"errors"

	"github.com/tsubasa66739/gin-nextjs-webapp/controller/schema"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/gorm"
)

type NoteService interface {
	GetNote(id uint) (model.TrnNote, error)
	CreateNote(req *schema.PostNoteReq) (model.TrnNote, error)
	UpdateNote(id uint, req *schema.PutNoteReq) error
}

type noteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(
	noteRepo repository.NoteRepository,
) NoteService {
	return &noteService{
		noteRepo: noteRepo,
	}
}

func (n *noteService) GetNote(id uint) (model.TrnNote, error) {
	note := model.TrnNote{}
	note.ID = &id
	err := n.noteRepo.GetById(&note)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return note, ErrNotFound
	}
	return note, err
}

func (n *noteService) CreateNote(req *schema.PostNoteReq) (model.TrnNote, error) {
	note := model.TrnNote{
		Title: req.Title,
		Body:  req.Body,
	}
	note.ID = nil
	err := n.noteRepo.Insert(&note)
	return note, err
}

func (n *noteService) UpdateNote(id uint, req *schema.PutNoteReq) error {
	note, err := n.GetNote(id)
	if err != nil {
		return err
	}

	req.Note.ID = &id
	if err = n.noteRepo.Update(&req.Note); err != nil {
		return err
	}

	return n.noteRepo.InsertHst(&note)
}