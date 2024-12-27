package service

import (
	"errors"

	"github.com/tsubasa66739/gin-nextjs-webapp/controller/schema"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/gorm"
)

type NoteService interface {
	GetNoteList() ([]model.TrnNote, error)
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

func (n *noteService) GetNoteList() ([]model.TrnNote, error) {
	notes := []model.TrnNote{}
	err := n.noteRepo.ListBy(&notes)
	return notes, err
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

	if err = n.noteRepo.InsertHst(note); err != nil {
		return err
	}

	note.Title = req.Title
	note.Body = req.Body
	return n.noteRepo.Update(&note)
}
