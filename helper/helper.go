package helper

import (
	"flashcardsProject/model"
	"time"

	"github.com/jinzhu/gorm"
)

func MakeFileNamePrefix() string {

	// ファイル名のプレフィックスを作成
	t := time.Now()
	layout := "2006_01_02_15_04_05"
	return t.Format(layout) + "_"

}

func CheckUser(userId string, password string) bool {

	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(create_user)")
	}
	defer db.Close()

	auth := false
	var user []model.User
	db.Find(&user)
	for _, w := range user {
		if userId == w.Name && password == w.Password {
			auth = true
			break
		}
	}
	return auth
}
