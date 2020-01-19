package model

import (
	"github.com/jinzhu/gorm"
)

/*
 * ユーザ全件取得
 */
func GetAllUser() []User {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(get_all_user)")
	}

	defer db.Close()

	var user []User
	db.Find(&user)
	return user

}

/*
 * ユーザ一件取得
 */
func GetUserDetail(id int) User {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(get_user_detail)")
	}
	defer db.Close()

	var user User
	db.First(&user, id)
	return user
}

/*
 * ユーザ作成
 */
func CreateUser(name string, password string) string {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(create_user)")
	}
	defer db.Close()

	isDouble := CountUser(name)

	msg := ""
	if isDouble == true {
		msg = "※ すでに登録されているユーザ名のため、登録できません"
		return msg
	}

	db.Create(&User{Name: name, Password: password})

	return msg
}

/*
 * ユーザの更新
 */
func UpdateUser(id int, name string, password string) {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(update_user)")
	}
	defer db.Close()

	var user User
	db.First(&user, id)
	user.Name = name
	user.Password = password
	db.Save(&user)
}

/*
 * ユーザの削除
 */
func DeleteUser(id int) {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database(delete_user")
	}

	defer db.Close()

	var user User
	db.First(&user, id)
	db.Delete(&user)

}

/*
 * 登録できるユーザ名か重複確認
 */
func CountUser(name string) bool {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(count_user)")
	}
	defer db.Close()
	var count int
	db.Debug().Model(&User{}).Where("name = ?", name).Count(&count)

	return count > 0
}
