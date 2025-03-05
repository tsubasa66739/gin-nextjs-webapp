package repository

import (
	"errors"

	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]model.Item, error)
	FindById(itemId uint) (*model.Item, error)
	Create(newItem model.Item) (*model.Item, error)
	Update(updateItem model.Item) (*model.Item, error)
	Delete(itemId uint) error
}

type ItemRepository struct {
	db *gorm.DB
}

type ItemMemoryRepository struct {
	items []model.Item
}

func NewItemMemoryRepository(items []model.Item) IItemRepository {
	return &ItemMemoryRepository{items: items} //interface型を戻すことで、実装がinterfaceの定義を満たしていない場合エラーが発生し未然に防げる
}

func (r *ItemMemoryRepository) FindAll() (*[]model.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*model.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (r *ItemMemoryRepository) Create(newItem model.Item) (*model.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(updateItem model.Item) (*model.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

func (r *ItemMemoryRepository) Delete(itemId uint) error {
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}

	return errors.New("Item not found")
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(newItem model.Item) (*model.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (r *ItemRepository) FindAll() (*[]model.Item, error) {
	var items []model.Item
	result := r.db.Find(&items) //検索結果をitemsに格納
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

// FindById implements IItemRepository.
func (r *ItemRepository) FindById(itemId uint) (*model.Item, error) {
	var item model.Item
	result := r.db.First(&item, itemId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

// Update implements IItemRepository.
func (r *ItemRepository) Update(updateItem model.Item) (*model.Item, error) {
	result := r.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

// Delete implements IItemRepository.
func (r *ItemRepository) Delete(itemId uint) error {
	deleteItem, err := r.FindById(itemId)
	if err != nil {
		return err
	}
	result := r.db.Delete(&deleteItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
