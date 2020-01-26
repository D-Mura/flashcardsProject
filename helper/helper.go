package helper

import (
	"flashcardsProject/config"
	"flashcardsProject/model"
	"time"

	"github.com/jinzhu/gorm"
)

/*
 * ファイル名のプレフィックス作成
 * アップロード日をアンダーバーつなぎのフォーマットにする
 */
func MakeFileNamePrefix() string {

	// ファイル名のプレフィックスを作成
	t := time.Now()
	layout := "2006_01_02_15_04_05"
	return t.Format(layout) + "_"

}

func PushedGoodButton(user string, id int) bool {
	goodUser := model.GetAllGood()

	for _, u := range goodUser {
		if u.GoodUser == user && u.WikiID == id {
			return true
		}
	}
	return false
}

func CheckUser(userId string, password string) bool {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))

	if err != nil {
		panic("failed to connect database(create_user)")
	}

	auth := false
	var user []model.User
	config.DB.Find(&user)
	for _, w := range user {
		if userId == w.Name && password == w.Password {
			auth = true
			break
		}
	}
	return auth
}
