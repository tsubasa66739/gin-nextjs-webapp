package main

import (
	"github.com/tsubasa66739/gin-nextjs-webapp/config"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Setup()
	db = repository.Setup()
}

func main() {
	server := controller.InitRouter(db)
	server.Run()
}
