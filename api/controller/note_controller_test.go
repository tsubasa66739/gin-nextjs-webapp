package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tsubasa66739/gin-nextjs-webapp/repository/model"
	"github.com/tsubasa66739/gin-nextjs-webapp/service"
	"go.uber.org/mock/gomock"
)

func TestGetNote(t *testing.T) {

	// モック作成
	mockCtrl := gomock.NewController(t)
	mockSvc := service.NewMockNoteService(mockCtrl)
	mockSvc.EXPECT().
		GetNote(gomock.Any()).
		DoAndReturn(func(id uint) (model.TrnNote, error) {
			ret := model.TrnNote{
				Title: "hoge",
				Body:  "fuga",
			}
			i := uint(1)
			ret.ID = &i
			ret.CreatedAt = time.Date(2024, 12, 31, 18, 45, 34, 0, time.Local)
			ret.UpdatedAt = time.Date(2025, 1, 5, 8, 25, 18, 0, time.Local)
			return ret, nil
		}).
		AnyTimes()

	// リクエスト定義
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/note", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.AddParam("id", "1")

	// テスト対象初期化
	noteCtrl := NewNoteController(mockSvc)

	// テスト実行
	noteCtrl.GetNote(c)

	var resBody map[string]interface{}
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatal(err)
	}

	if res.Code != http.StatusOK {
		t.Errorf("TestGetNote, want: 200, got: %d", res.Code)
	}
	if resBody["title"] != "hoge" {
		t.Errorf("TestGetNote, want: hoge, got: %s", resBody["title"])
	}
	if resBody["body"] != "fuga" {
		t.Errorf("TestGetNote, want: fuga, got: %s", resBody["body"])
	}

}
