package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Body struct {
}

type Wiki struct {
}

type User struct {
	gorm.Model
	Name     string
	Password string
}

/*
 * DB初期化処理
 */
func db_init() {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		fmt.Println("failed to connect database(init)")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

/*
 * ユーザ一件取得
 */
func get_user_detail(id int) User {
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
 * ユーザ全件取得
 */
func get_all_user() []User {
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
 * ユーザ作成
 */
func create_user(name string, password string) {
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
func update_user(id int, name string, password string) {
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
func delete_user(id int) {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(delete_user")
	}

	defer db.Close()

	var user User
	db.First(&user, id)
	db.Delete(&user)

}

func main() {
	r := gin.Default()

	// 静的ファイルの読み込み
	r.Static("/assets", "./assets")

	// Templateの読み込み
	r.LoadHTMLGlob("./templates/*")

	db_init()

	// ユーザ情報全件取得
	r.GET("/user", func(c *gin.Context) {
		user := get_all_user()

		c.HTML(200, "userInfo.tmpl", gin.H{
			"user": user,
		})
	})

	// ユーザ情報を一件取得
	r.GET("/user/:id", func(c *gin.Context) {
		num := c.Param("id")
		id, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		user := get_user_detail(id)
		c.HTML(200, "userInfoDetail.tmpl", gin.H{
			"user": user,
		})
	})

	// ユーザ新規作成
	r.POST("/new_user", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		create_user(name, password)
		c.Redirect(302, "/user")
	})

	// ユーザの更新
	r.POST("/user/:id/update", func(c *gin.Context) {
		num := c.Param("id")
		id, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		name := c.PostForm("name")
		password := c.PostForm("password")

		update_user(id, name, password)
		c.Redirect(302, "/user/"+num)

	})

	// ユーザの削除
	r.POST("/user/:id/delete", func(c *gin.Context) {
		num := c.Param("id")
		id, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		delete_user(id)
		c.Redirect(302, "/user")
	})

	r.Run()
}
