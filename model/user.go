package model

import (
	"flashcardsProject/config"

	"github.com/jinzhu/gorm"
)

/*
 * ユーザ全件取得
 */
func GetAllUser() []User {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(get_all_user)")
	}

	defer config.DB.Close()

	var user []User
	config.DB.Find(&user)
	return user

}

/*
 * ユーザ一件取得
 */
func GetUserDetail(id int) User {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(get_user_detail)")
	}
	defer config.DB.Close()

	var user User
	config.DB.First(&user, id)
	return user
}

/*
 * ユーザ作成
 */
func CreateUser(name string, password string) string {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(create_user)")
	}
	defer config.DB.Close()

	isDouble := CountUser(name)

	msg := ""
	if isDouble == true {
		msg = "※ すでに登録されているユーザ名のため、登録できません"
		return msg
	}

	config.DB.Create(&User{Name: name, Password: password})

	return msg
}

/*
 * ユーザの更新
 */
func UpdateUser(id int, name string, password string) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(update_user)")
	}
	defer config.DB.Close()

	var user User
	config.DB.First(&user, id)
	user.Name = name
	user.Password = password
	config.DB.Save(&user)
}

/*
 * ユーザの削除
 */
func DeleteUser(id int) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(delete_user")
	}

	defer config.DB.Close()

	var user User
	config.DB.First(&user, id)
	config.DB.Delete(&user)

}

/*
 * 登録できるユーザ名か重複確認
 */
func CountUser(name string) bool {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(count_user)")
	}
	defer config.DB.Close()
	var count int
	config.DB.Debug().Model(&User{}).Where("name = ?", name).Count(&count)

	return count > 0
}
