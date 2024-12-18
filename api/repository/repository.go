package repository

import (
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
