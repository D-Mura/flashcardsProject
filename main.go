package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 静的ファイルの読み込み
	r.Static("/assets", "./assets")

	// Templateの読み込み
	r.LoadHTMLGlob("./templates/*")

	//db_init()

	// 全件取得
	r.GET("/", func(c *gin.Context) {
		//people := get_all()
		c.HTML(200, "index.tmpl", gin.H{
			"people": "people",
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
