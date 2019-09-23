package main

import (
	"fmt"

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
	Name string
}

func db_init() {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		fmt.Println("failed to connect database(init)")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func get_all_User() []User {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(get_all)")
	}

	defer db.Close()

	var user []User
	db.Find(&user)
	return user

}

func main() {
	r := gin.Default()

	// 静的ファイルの読み込み
	r.Static("/assets", "./assets")

	// Templateの読み込み
	r.LoadHTMLGlob("./templates/*")

	db_init()

	// 全件取得
	r.GET("/", func(c *gin.Context) {
		user := get_all_User()

		c.HTML(200, "index.tmpl", gin.H{
			"user": user,
		})
	})
	/*

		// 新規作成
		r.POST("/new", func(c *gin.Context) {
			name := c.PostForm("name")
			age, _ := strconv.Atoi(c.PostForm("age"))
			create(name, age)
			c.Redirect(302, "/")
		})
	*/
	r.Run()
}
