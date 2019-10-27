package main

import (
	"flashcardsProject/controller"
	"flashcardsProject/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
 * DB初期化処理
 */
func db_init() {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("failed to connect database(init)")
	}
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Wiki{}, &model.Body{}, &model.User{})

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
	r.GET("/user", controller.GetAllUser)

	// ユーザ情報を一件取得
	r.GET("/user/:id", controller.GetUserDetail)

	// ユーザ新規作成
	r.POST("/new_user", controller.CreateUser)

	// ユーザの更新
	r.POST("/user/:id/update", controller.UpdateUser)

	// ユーザの削除
	r.POST("/user/:id/delete", controller.DeleteUser)

	// Wiki全件取得
	r.GET("/wiki", controller.GetAllWiki)

	// Wikiのソート
	r.GET("/sort_wiki", controller.SortAllWiki)

	// Wiki内容取得
	r.GET("/wiki/:id", controller.GetWikiDetail)

	// Wiki新規作成
	r.GET("/new_wiki", controller.NewWiki)

	// Wiki作成
	r.POST("/create_wiki", controller.CreateWiki)

	// Wikiの更新
	// 画像の更新は非対応
	r.POST("/update_wiki/:id", controller.UpdateWiki)

	// Wikiの削除
	r.POST("/delete_wiki/:id", controller.DeleteWiki)

	// Wikiの検索
	r.POST("/search_wiki", controller.SearchWiki)

	// Wiki画像の更新
	r.POST("/update_wiki_picture", controller.UpdateWikiPicture)

	/*
		// FileUpload Test
		r.POST("/upload", controller.Upload)
	*/

	r.Run()
}
