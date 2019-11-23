package controller

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	//セッションからデータを破棄する
	session := sessions.Default(c)
	log.Println("セッション取得")
	session.Clear()
	log.Println("クリア処理")
	session.Save()

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
