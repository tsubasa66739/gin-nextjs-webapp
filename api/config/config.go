package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Setup() {
	loadEnv()
}

// .envファイルを読み込む
func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading env target", err)
	}
}
