package main

import (
	"github.com/tsubasa66739/gin-nextjs-webapp/config"
	"github.com/tsubasa66739/gin-nextjs-webapp/controller"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository"
)

func init() {
	config.Setup()
	repository.Setup()
}

func main() {
	server := controller.InitRouter()
	server.Run()
}
