package controller

import (
	"flashcardsProject/helper"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, UserId string) {

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
	if helper.CheckUser(UserId, Password) == false {
		GetLogout(c)
	} else {
		Login(c, UserId) // // 同じパッケージ内のログイン処理

		c.Redirect(http.StatusMovedPermanently, "/wiki")
	}

}
