package repository

import (
	"errors"
	"fmt"
	"os"

	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormSchema "gorm.io/gorm/schema"
)

func Setup() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	dbConfig := &gorm.Config{
		NamingStrategy: gormSchema.NamingStrategy{
			TablePrefix:   "tb_", // テーブル名のPrefix
			SingularTable: true,  // テーブル名を複数形にしない
		},
	}
	db, err := gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(
		&model.TrnNote{},
		&model.HstNote{},
	)
	return db
}

type IItemRepository interface {
	FindAll() (*[]model.Item, error)
	FindById(itemId uint) (*model.Item, error)
	Create(newItem model.Item) (*model.Item, error)
	Update(updateItem model.Item) (*model.Item, error)
	Delete(itemId uint) error
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
