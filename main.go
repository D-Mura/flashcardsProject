package main

import (
	"flashcardsProject/controller"
	"flashcardsProject/model"
	"fmt"
	"log"
	"net/http"

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

	r.POST("/login", PostLogin)

	r.Use(sessionCheck())

	r.POST("/logout", PostLogout)

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

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("aaaaa")
		session := sessions.Default(c)
		LoginInfo.UserId = session.Get("UserId")
		log.Println(session.Get("UserId"))

		// セッションがない場合、ログインフォームをだす
		if LoginInfo.UserId == nil {
			log.Println("ログインしていません")
			//c.Redirect(http.StatusMovedPermanently, "/wiki")
			//c.Abort() // これがないと続けて処理されてしまう
		} else {
			log.Println("cc")

			c.Set("UserId", LoginInfo.UserId) // ユーザidをセット
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}

func Logout(c *gin.Context) {

	//セッションからデータを破棄する
	session := sessions.Default(c)
	log.Println("セッション取得")
	session.Clear()
	log.Println("クリア処理")
	session.Save()

}

func Login(c *gin.Context, UserId string) {
	log.Println("bbb")

	//セッションにデータを格納する
	session := sessions.Default(c)
	session.Set("UserId", UserId)
	session.Save()
}

func PostLogin(c *gin.Context) {
	log.Println("ログイン処理")
	UserId := c.PostForm("userId")
	Password := c.PostForm("password")

	log.Println(UserId + " " + Password)
	// login check
	if UserId != "test" || Password != "test" {
		GetLogout(c)
	} else {
		Login(c, UserId) // // 同じパッケージ内のログイン処理

		c.Redirect(http.StatusMovedPermanently, "/wiki")
	}

}

func PostLogout(c *gin.Context) {
	log.Println("ログアウト処理")
	Logout(c) // 同じパッケージ内のログアウト処理

	c.HTML(200, "login.tmpl", gin.H{
		"Error": "",
	})
}

func GetLogout(c *gin.Context) {
	log.Println("ログアウト処理")
	Logout(c) // 同じパッケージ内のログアウト処理

	c.HTML(200, "login.tmpl", gin.H{
		"Error": "error",
	})
}
