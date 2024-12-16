package service

import (
	"errors"
	"fmt"

	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/schema"
	"gorm.io/gorm"
)

func GetNote(id uint) (repository.TrnNote, error) {
	note := repository.TrnNote{}
	note.ID = &id
	err := repository.GetNoteById(&note)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return note, ErrNotFound
	}
	return note, err
}

func CreateNote(req *schema.PostNoteReq) (repository.TrnNote, error) {
	note := repository.TrnNote{
		Title: req.Title,
		Body:  req.Body,
	}
	note.ID = nil
	err := repository.InsertNote(&note)
	return note, err
}

func UpdateNote(id uint, req *schema.PutNoteReq) error {
	note, err := GetNote(id)
	if err != nil {
		return err
	}

	req.Note.ID = &id
	fmt.Println(req.Note)
	if err = repository.UpdateNote(&req.Note); err != nil {
		return err
	}

	return repository.InsertNoteHst(&note)
}
