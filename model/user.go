package model

import "github.com/jinzhu/gorm"

/*
 * ユーザ全件取得
 */
func GetAllUser() []User {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
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
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
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
func CreateUser(name string, password string) {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(create_user)")
	}
	defer db.Close()

	db.Create(&User{Name: name, Password: password})
}

/*
 * ユーザの更新
 */
func UpdateUser(id int, name string, password string) {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
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
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(delete_user")
	}

	defer db.Close()

	var user User
	db.First(&user, id)
	db.Delete(&user)

}
