package config

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	loadEnv()
}

// .envファイルを読み込む
func loadEnv() {
	err := godotenv.Load(".env") //.envファイルから環境変数を読み込む
	if err != nil {
		log.Fatal("Error loading env target", err)
	}
}
