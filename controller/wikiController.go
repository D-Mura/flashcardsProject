package controller

import (
	"flashcardsProject/helper"
	"flashcardsProject/model"
	"log"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
	imageupload "github.com/olahol/go-imageupload"
)

// Wiki全件取得
func GetAllWiki(c *gin.Context) {
	UserId, _ := c.Get("UserId")
	bySorted := "title"
	wikiForScreenA, wikiForScreenB, wikiForScreenC := model.GetAllWiki(bySorted)

	isSearched := false

	c.HTML(200, "index.tmpl", gin.H{
		"wikiForScreenA": wikiForScreenA,
		"wikiForScreenB": wikiForScreenB,
		"wikiForScreenC": wikiForScreenC,
		"isSearched":     isSearched,
		"UserId":         UserId,
	})
}

// Wikiのソート
func SortAllWiki(c *gin.Context) {
	UserId, _ := c.Get("UserId")

	bySorted := c.Query("sort-btn")
	wikiForScreenA, wikiForScreenB, wikiForScreenC := model.GetAllWiki(bySorted)

	isSearched := false
	c.HTML(200, "index.tmpl", gin.H{
		"wikiForScreenA": wikiForScreenA,
		"wikiForScreenB": wikiForScreenB,
		"wikiForScreenC": wikiForScreenC,
		"isSearched":     isSearched,
		"UserId":         UserId,
	})
}

// Wiki内容取得
func GetWikiDetail(c *gin.Context) {
	UserId, _ := c.Get("UserId")

	num := c.Param("id")
	id, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	wiki := model.GetWikiDetail(id)

	c.HTML(200, "detail.tmpl", gin.H{
		"wiki":   wiki,
		"UserId": UserId,
	})
}

// Wikiの削除
func DeleteWiki(c *gin.Context) {

	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	model.DeleteWiki(id)
	c.Redirect(302, "/wiki")

}

// Wiki新規作成
func NewWiki(c *gin.Context) {
	c.HTML(200, "new.tmpl", gin.H{})
}

// Wiki作成
func CreateWiki(c *gin.Context) {
	title := c.PostForm("title")
	num := c.PostForm("screenId")
	author := c.PostForm("author")
	text := c.PostForm("text")
	url := c.PostForm("url")

	screenId, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	pictName := "default.png"
	// ファイルのアップロード処理
	img, err := imageupload.Process(c.Request, "file")
	if err == nil {
		log.Println(img.Filename)

		// 300x300にリサイズ
		thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
		if err != nil {
			panic(err)
		}

		// ファイル名のプレフィックスを作成
		pictName = helper.MakeFileNamePrefix() + img.Filename

		// ファイルの保存
		thumb.Save("./assets/image/" + pictName)
	}

	body := model.Body{Text: text, Author: author, Url: url}
	wiki := model.Wiki{Title: title, PictureName: pictName, ScreenID: screenId, Body: body}
	model.CreateWiki(wiki)
	c.Redirect(302, "/wiki")

}

// Wikiの更新
func UpdateWiki(c *gin.Context) {
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
	nBody := model.Body{Text: text, Author: author, Url: url}
	nWiki := model.Wiki{Title: title, ScreenID: screenId, Body: nBody}

	model.UpdateWiki(id, nWiki)
	c.Redirect(302, "/wiki")

}

// SearchWiki
func SearchWiki(c *gin.Context) {
	UserId, _ := c.Get("UserId")

	word := c.PostForm("word")
	option := c.PostForm("search-option")
	log.Println(option)
	wikiForScreenA, wikiForScreenB, wikiForScreenC := model.SearchWiki(word, option)

	isSearched := true
	c.HTML(200, "index.tmpl", gin.H{
		"wikiForScreenA": wikiForScreenA,
		"wikiForScreenB": wikiForScreenB,
		"wikiForScreenC": wikiForScreenC,
		"isSearched":     isSearched,
		"UserId":         UserId,
	})

}

// UpdateGood
func UpdateGood(c *gin.Context) {

	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	model.UpdateGood(id)
	c.Redirect(302, "/wiki/"+n)
}

// UpdateWikiPicture
func UpdateWikiPicture(c *gin.Context) {
	n := c.PostForm("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	// ファイルのアップロード処理
	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		panic(err)
	}
	// 300x300にリサイズ
	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
	if err != nil {
		panic(err)
	}

	// ファイル名のプレフィックスを作成
	pictName := helper.MakeFileNamePrefix() + img.Filename

	// ファイルの保存
	thumb.Save("./assets/image/" + pictName)

	model.UpdateWikiPicture(id, pictName)
	c.Redirect(302, "/wiki")

}

/*
// FileUpload
func Upload(c *gin.Context) {
	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		panic(err)
	}
	log.Println(img.Filename)
	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
	if err != nil {
		panic(err)
	}
	thumb.Save("./assets/image/" + img.Filename)
	c.Redirect(302, "/wiki")
}
*/
