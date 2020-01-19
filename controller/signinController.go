package controller

import (
	"flashcardsProject/model"
	"log"

	"github.com/gin-gonic/gin"
)

// SignIn画面の表示
func GetSignIn(c *gin.Context) {
	c.HTML(200, "signin.tmpl", gin.H{})
}

// User登録
func PostSignIn(c *gin.Context) {
	log.Println("aaaa")
	name := c.PostForm("name")
	password := c.PostForm("password")
	message := model.CreateUser(name, password)

	if message != "" {
		c.HTML(200, "signin.tmpl", gin.H{
			"message": message,
		})
	}
	c.Redirect(302, "/wiki")
}
