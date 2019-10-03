package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

/*
 * Wiki一覧取得
 */
func GetAllWiki() []Wiki {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(GetAllWiki())")
	}
	defer db.Close()

	var wiki []Wiki
	db.Find(&wiki)
	return wiki

}

/*
 * Wiki詳細取得
 */
func GetWikiDetail(id int) Wiki {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(GetWikiDetail())")
	}
	defer db.Close()

	var wiki Wiki
	var body Body

	db.Debug().Find(&wiki, id).Related(&body)

	// 上記のSQLだと、wiki.Bodyに入らないので
	// 代入する
	wiki.Body = body

	log.Println(wiki.Body)
	return wiki
}

/*
 * Wikiの削除
 */
func DeleteWiki(id int) {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(update_wiki)")
	}
	defer db.Close()

	var wiki Wiki
	db.First(&wiki, id)
	db.Delete(&wiki)
}

/*
 * Wikiの新規作成
 */
func CreateWiki(wiki Wiki) {
	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(create_wiki)")
	}

	defer db.Close()
	db.Debug().Create(&wiki)
}

/*
 * Wikiの更新
 */
func UpdateWiki(id int, nWiki Wiki) {

	db, err := gorm.Open("sqlite3", "testGin.sqlite3")
	if err != nil {
		panic("failed to connect database(update_wiki)")
	}
	defer db.Close()

	var wiki Wiki
	db.First(&wiki, id)
	wiki.Title = nWiki.Title
	wiki.ScreenID = nWiki.ScreenID
	wiki.Body = nWiki.Body
	db.Debug().Save(&wiki)

}
