package model

import (
	"flashcardsProject/config"
	"log"

	"github.com/jinzhu/gorm"
)

/*
 * Wiki一覧取得
 */
func GetAllWiki(bySorted string) ([]Wiki, []Wiki, []Wiki) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(GetAllWiki())")
	}

	var wiki []Wiki
	if bySorted == "date" {
		config.DB.Debug().Order("updated_at").Find(&wiki)
	} else {
		config.DB.Debug().Order("title").Find(&wiki)

	}

	var wikiForScreenA, wikiForScreenB, wikiForScreenC []Wiki

	/*
	 * ToDo: switch文に変更
	 */

	// 画面IDに応じて振り分ける
	for _, w := range wiki {
		if w.ScreenID == 1 {
			wikiForScreenA = append(wikiForScreenA, w)
		} else if w.ScreenID == 2 {
			wikiForScreenB = append(wikiForScreenB, w)
		} else if w.ScreenID == 3 {
			wikiForScreenC = append(wikiForScreenC, w)
		}

		if w.PictureName == "" {

			w.PictureName = "default.png"
		}

	}

	return wikiForScreenA, wikiForScreenB, wikiForScreenC

}

/*
 * Wiki詳細取得
 */
func GetWikiDetail(id int) Wiki {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(GetWikiDetail())")
	}

	var wiki Wiki
	var body Body

	// BodyのPVを閲覧に応じて加算する
	UpdatePageView(id)

	config.DB.Debug().Find(&wiki, id).Related(&body)

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
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(update_wiki)")
	}

	var wiki Wiki
	config.DB.First(&wiki, id)
	config.DB.Delete(&wiki)
}

/*
 * Wikiの新規作成
 */
func CreateWiki(wiki Wiki) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(create_wiki)")
	}

	config.DB.Debug().Create(&wiki)
}

/*
 * Wikiの更新
 */
func UpdateWiki(id int, nWiki Wiki) {

	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(update_wiki)")
	}

	var wiki Wiki
	config.DB.First(&wiki, id)
	wiki.Title = nWiki.Title
	wiki.ScreenID = nWiki.ScreenID
	wiki.Body = nWiki.Body
	config.DB.Debug().Save(&wiki)

}

/*
 * Wikiの検索
 */
func SearchWiki(word string, option string) ([]Wiki, []Wiki, []Wiki) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(serach_wiki)")
	}
	var wiki []Wiki

	// 全文検索: full-search
	// 部分検索: partial-search
	if option == "full-search" {
		config.DB.Debug().Where("title = ?", word).Find(&wiki)
	} else if option == "partial-search" {
		// SQLite3の場合、GLOBを用いる
		// PostgresやMySQLでは異なるので注意
		// ex-postgres)  "title ~ ?", "^"+word+"^"のようになると思われる
		config.DB.Debug().Where("title LIKE ?", "%"+word+"%").Find(&wiki)
	}

	var wikiForScreenA, wikiForScreenB, wikiForScreenC []Wiki

	/*
	 * ToDo: switch文に変更
	 */

	// 画面IDに応じて振り分ける
	for _, w := range wiki {
		if w.ScreenID == 1 {
			wikiForScreenA = append(wikiForScreenA, w)
		} else if w.ScreenID == 2 {
			wikiForScreenB = append(wikiForScreenB, w)
		} else if w.ScreenID == 3 {
			wikiForScreenC = append(wikiForScreenC, w)
		}
	}

	return wikiForScreenA, wikiForScreenB, wikiForScreenC

}

/*
 * Good数の更新(+1)
 */
func UpdateGood(id int) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(update_good)")
	}

	var wiki Wiki
	config.DB.First(&wiki, id)
	wiki.Good += 1
	config.DB.Debug().Save(&wiki)
}

/*
 * Wiki画像の更新
 */
func UpdateWikiPicture(id int, pictName string) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(update_wiki_picture)")
	}

	var wiki Wiki
	config.DB.First(&wiki, id)
	wiki.PictureName = pictName
	config.DB.Debug().Save(&wiki)
}
