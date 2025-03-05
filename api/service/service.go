package service

import (
	"errors"

	"github.com/tsubasa66739/gin-nextjs-webapp/dto"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
)

var (
	ErrNotFound       = errors.New("resource not found")
	ErrInternalServer = errors.New("unknown error")
)

type IItemService interface {
	FindAll() (*[]model.Item, error)
	FindById(itemId uint) (*model.Item, error)
	Create(createItemInput dto.CreateItemInput) (*model.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*model.Item, error)
	Delete(itemId uint) error
}

type ItemService struct {
	repository repository.IItemRepository
}

func NewItemService(repository repository.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]model.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*model.Item, error) {
	return s.repository.FindById(itemId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*model.Item, error) {
	newItem := model.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		Soldout:     false, //出品時は必ずfalse
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*model.Item, error) {
	targetItem, err := s.FindById(itemId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}

	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}

	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}

	if updateItemInput.Soldout != nil {
		targetItem.Soldout = *updateItemInput.Soldout
	}

	return s.repository.Update(*targetItem)

}

func (s *ItemService) Delete(itemId uint) error {
	return s.repository.Delete(itemId)
}
