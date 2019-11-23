package main

import (
	"flashcardsProject/controller"
	"flashcardsProject/model"
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var LoginInfo model.SessionInfo

/*
 * DB初期化処理
 */
func db_init() {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
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

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// 静的ファイルの読み込み
	r.Static("/assets", "./assets")

	// Templateの読み込み
	r.LoadHTMLGlob("./templates/*")

	db_init()

	r.POST("/login", controller.PostLogin)

	r.Use(sessionCheck())

	r.POST("/logout", controller.PostLogout)
	/*
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
	*/
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

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		LoginInfo.UserId = session.Get("UserId")
		log.Println(session.Get("UserId"))

		// セッションがない場合、ログインフォームをだす
		if LoginInfo.UserId == nil {
			log.Println("ログインしていません")
			//c.Redirect(http.StatusMovedPermanently, "/wiki")
			//c.Abort() // これがないと続けて処理されてしまう
		} else {

			c.Set("UserId", LoginInfo.UserId) // ユーザidをセット
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}
