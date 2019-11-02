package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

/*
 * Wiki一覧取得
 */
func GetAllWiki(bySorted string) ([]Wiki, []Wiki, []Wiki) {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(GetAllWiki())")
	}
	defer db.Close()

	var wiki []Wiki
	if bySorted == "date" {
		db.Debug().Order("updated_at").Find(&wiki)
	} else {
		db.Debug().Order("title").Find(&wiki)

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
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
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
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
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
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
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

	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
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

/*
 * Wikiの検索
 */
func SearchWiki(word string, option string) ([]Wiki, []Wiki, []Wiki) {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(serach_wiki)")
	}
	var wiki []Wiki

	// 全文検索: full-search
	// 部分検索: partial-search
	if option == "full-search" {
		db.Debug().Where("title = ?", word).Find(&wiki)
	} else if option == "partial-search" {
		// SQLite3の場合、GLOBを用いる
		// PostgresやMySQLでは異なるので注意
		// ex-postgres)  "title ~ ?", "^"+word+"^"のようになると思われる
		db.Debug().Where("title LIKE ?", "%"+word+"%").Find(&wiki)
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
 * Wiki画像の更新
 */
func UpdateWikiPicture(id int, pictName string) {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(update_wiki_picture)")
	}
	defer db.Close()

	var wiki Wiki
	db.First(&wiki, id)
	wiki.PictureName = pictName
	db.Debug().Save(&wiki)
}
