package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Wiki struct {
	gorm.Model
	Body     Body
	Title    string
	ScreenID int
}

type Body struct {
	gorm.Model
	Text   string
	Author string
	Url    string
	WikiID int `gorm:"index"`
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

	db.AutoMigrate(&Wiki{}, &Body{}, &User{})
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

/*
 * Wiki一覧取得
 */
func get_all_wiki() []Wiki {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(get_all_wiki)")
	}
	defer db.Close()

	var wiki []Wiki
	db.Find(&wiki)
	return wiki

}

/*
 * Wiki詳細取得
 */
func get_wiki_detail(id int) Wiki {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(get_wiki_detail)")
	}
	defer db.Close()

	var wiki Wiki
	var body Body

	db.Debug().Find(&wiki, id).Related(&body)

	// 上記のSQLだと、wiki.Bodyに入らないので
	// 代入する
	wiki.Body = body

	log.Println(wiki.Body)
	return wiki
}

/*
 * Wikiの新規作成
 */
func create_wiki(wiki Wiki) {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(create_wiki)")
	}

	defer db.Close()
	db.Debug().Create(&wiki)
}

/*
 * Wikiの更新
 */
func update_wiki(id int, nWiki Wiki) {

	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(update_wiki)")
	}
	defer db.Close()

	var wiki Wiki
	db.First(&wiki, id)
	wiki.Title = nWiki.Title
	wiki.ScreenID = nWiki.ScreenID
	wiki.Body = nWiki.Body
	db.Debug().Save(&wiki)

}

/*
 * Wikiの削除
 */
func delete_wiki(id int) {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(update_wiki)")
	}
	defer db.Close()

	var wiki Wiki
	db.First(&wiki, id)
	db.Delete(&wiki)
}

/*
 * Main
 */
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

	// Wiki全件取得
	r.GET("/wiki", func(c *gin.Context) {
		wiki := get_all_wiki()

		c.HTML(200, "index.tmpl", gin.H{
			"wiki": wiki,
		})
	})

	// Wiki内容取得
	r.GET("/wiki/:id", func(c *gin.Context) {
		num := c.Param("id")
		id, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		wiki := get_wiki_detail(id)

		c.HTML(200, "detail.tmpl", gin.H{
			"wiki": wiki,
		})
	})

	// Wiki新規作成
	r.GET("/new_wiki", func(c *gin.Context) {
		c.HTML(200, "new.tmpl", gin.H{})
	})

	// Wiki作成
	r.POST("/create_wiki", func(c *gin.Context) {
		title := c.PostForm("title")
		num := c.PostForm("screenId")
		author := c.PostForm("author")
		text := c.PostForm("text")
		url := c.PostForm("url")

		screenId, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		body := Body{Text: text, Author: author, Url: url}
		wiki := Wiki{Title: title, ScreenID: screenId, Body: body}
		create_wiki(wiki)
		c.Redirect(302, "/wiki")

	})

	// Wikiの更新
	r.POST("/update_wiki/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		title := c.PostForm("title")
		num := c.PostForm("screenId")
		author := c.PostForm("author")
		text := c.PostForm("text")
		url := c.PostForm("url")

		screenId, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		nBody := Body{Text: text, Author: author, Url: url}
		nWiki := Wiki{Title: title, ScreenID: screenId, Body: nBody}

		update_wiki(id, nWiki)
		c.Redirect(302, "/wiki")

	})

	r.POST("/delete_wiki/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		delete_wiki(id)
		c.Redirect(302, "/wiki")

	})

	r.Run()
}
