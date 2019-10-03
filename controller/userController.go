package controller

import (
	"flashcardsProject/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ユーザ情報全件取得
func GetAllUser(c *gin.Context) {
	user := model.GetAllUser()

	c.HTML(200, "userInfo.tmpl", gin.H{
		"user": user,
	})
}

// ユーザ情報を一件取得
func GetUserDetail(c *gin.Context) {
	num := c.Param("id")
	id, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	user := model.GetUserDetail(id)
	c.HTML(200, "userInfoDetail.tmpl", gin.H{
		"user": user,
	})
}

// ユーザ新規作成
func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	model.CreateUser(name, password)
	c.Redirect(302, "/user")
}

func UpdateUser(c *gin.Context) {
	num := c.Param("id")
	id, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	name := c.PostForm("name")
	password := c.PostForm("password")

	model.UpdateUser(id, name, password)
	c.Redirect(302, "/user/"+num)

}

// ユーザの削除
func DeleteUser(c *gin.Context) {
	num := c.Param("id")
	id, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	model.DeleteUser(id)
	c.Redirect(302, "/user")
}
