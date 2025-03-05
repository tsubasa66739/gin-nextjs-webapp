package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// カスタム型JSONBを定義
type JSONB map[string]interface{}

// データベースから読み込んだ値をJSONB型に変換するためのメソッド
func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	//[]byteをJSONB型にデコード
	return json.Unmarshal(bytes, j)
}

// JSONB型の値をデータベースに保存するために[]byteに変換するメソッド
func (j JSONB) Value() (driver.Value, error) {

	//JSONB型をJSON形式の[]byteにエンコード
	return json.Marshal(j)
}
