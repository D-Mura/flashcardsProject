package controller

import (
	"flashcardsProject/model"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	imageupload "github.com/olahol/go-imageupload"
)

// Wiki全件取得
func GetAllWiki(c *gin.Context) {
	wikiForScreenA, wikiForScreenB, wikiForScreenC := model.GetAllWiki()

	c.HTML(200, "index.tmpl", gin.H{
		"wikiForScreenA": wikiForScreenA,
		"wikiForScreenB": wikiForScreenB,
		"wikiForScreenC": wikiForScreenC,
	})
}

// Wiki内容取得
func GetWikiDetail(c *gin.Context) {
	num := c.Param("id")
	id, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	wiki := model.GetWikiDetail(id)

	c.HTML(200, "detail.tmpl", gin.H{
		"wiki": wiki,
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

	// ファイルのアップロード処理
	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		panic(err)
	}
	log.Println(img.Filename)

	// 300x300にリサイズ
	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
	if err != nil {
		panic(err)
	}

	// ファイル名のプレフィックスを作成
	/*
	 * 後ほど共通化（日付処理）
	 */
	t := time.Now()
	layout := "2006_01_02_15_04_05"
	pictName := t.Format(layout) + "_" + img.Filename

	// ファイルの保存
	thumb.Save("./assets/image/" + pictName)

	body := model.Body{Text: text, Author: author, Url: url}
	wiki := model.Wiki{Title: title, PictureName: pictName, ScreenID: screenId, Body: body}
	model.CreateWiki(wiki)
	c.Redirect(302, "/wiki")

}

// Wikiの更新
// 画像の更新は非対応
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
