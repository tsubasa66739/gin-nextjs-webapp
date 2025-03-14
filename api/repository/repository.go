package repository

import (
	// "errors"
	"fmt"
	"os"

	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm" //ORMライブラリはDB操作を簡単に（SQLクエリを書かない）するツール
	gormSchema "gorm.io/gorm/schema"
)

func Setup() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", //環境変数を取得し、それぞれ代入
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	dbConfig := &gorm.Config{ //GORMの設定を指定
		NamingStrategy: gormSchema.NamingStrategy{
			TablePrefix:   "tb_", // テーブル名のPrefix
			SingularTable: true,  // テーブル名を複数形にしない
		},
	}
	db, err := gorm.Open(postgres.Open(dsn), dbConfig) //PostgreSQLに接続しdsnをひらく、GORMの設定を読み込む
	if err != nil {
		panic(err.Error()) //プログラムの実行が即座に停止、リカバリー不可能なエラーのハンドリングを行う
	}
	db.AutoMigrate( //gormを使用してDBにそれぞれのテーブルを作成
		&model.TrnNote{}, //それぞれのモデルに基づいてテーブル作成
		&model.HstNote{},
		&model.Item{},
		&model.User{},
	)
	return db
}
