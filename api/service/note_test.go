package service

import (
	"errors"
	"testing"

	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetNote(t *testing.T) {

	// モックの準備
	mockCtrl := gomock.NewController(t)
	mockRepo := repository.NewMockNoteRepository(mockCtrl)

	// モック定義
	mockRepo.EXPECT().
		GetById(gomock.Any()).
		DoAndReturn(func(note *model.TrnNote) error {
			*note.ID = 2
			note.Title = "this is the note."
			note.Body = "hope to pass the test."
			return nil
		}).
		AnyTimes()

	// テスト対象初期化
	noteService := NewNoteService(mockRepo)

	// テスト実行
	note, err := noteService.GetNote(2)
	if err != nil {
		t.Fatal(err)
	}
	if *note.ID != 2 {
		t.Errorf("TestGetNote want: 2, got: %d", *note.ID)
	}
}

func TestGetNote_NotFound(t *testing.T) {

	// モックの準備
	mockCtrl := gomock.NewController(t)
	mockRepo := repository.NewMockNoteRepository(mockCtrl)

	// モック定義
	mockRepo.EXPECT().
		GetById(gomock.Any()).
		DoAndReturn(func(note *model.TrnNote) error {
			return gorm.ErrRecordNotFound
		}).
		AnyTimes()

	// テスト対象初期化
	noteService := NewNoteService(mockRepo)

	// テスト実行
	_, err := noteService.GetNote(2)
	if err == nil {
		t.Fail()
	} else if !errors.Is(ErrNotFound, err) {
		t.Errorf("TestGetNote_NotFound: want: service.ErrNotFound, got: %#v", err)
	}
}
