package service

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/tsubasa66739/gin-nextjs-webapp/controller/schema"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	//"gorm.io/gorm"
)

type IItemService interface {
	FindAll() (*[]model.Item, error)
	FindById(itemId uint) (*model.Item, error)
	Create(CreateItemInput *schema.CreateItemInput) (*model.Item, error)
	Update(itemId uint, updateItemInput *schema.UpdateItemInput) (*model.Item, error)
	Delete(itemId uint) error
	ExportItemToCSV(itemCSV string) error
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
func (s *ItemService) Create(CreateItemInput *schema.CreateItemInput) (*model.Item, error) {
	newItem := model.Item{
		Name:        CreateItemInput.Name,
		Price:       CreateItemInput.Price,
		Description: CreateItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)

}
func (s *ItemService) Update(itemId uint, updateItemInput *schema.UpdateItemInput) (*model.Item, error) {
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
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return s.repository.Update(*targetItem)
}

func (s *ItemService) Delete(itemId uint) error {
	return s.repository.Delete(itemId)
}

// CSV出力
func (s *ItemService) ExportItemToCSV(itemCSV string) error {
	items, err := s.repository.FindAll()
	if err != nil {
		return nil
	}
	file, err := os.Create(itemCSV)
	if err != nil {
		return nil
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"ID", "Name", "Price", "Description"})
	for _, item := range *items {
		writer.Write([]string{
			fmt.Sprintf("%d", item.ID),
			item.Name,
			fmt.Sprintf("%d", item.Price),
			item.Description,
		})
	}
	return nil
}
