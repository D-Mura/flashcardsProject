package model

import (
	"github.com/jinzhu/gorm"
)

/*
 * Good数の更新(+1)
 */
func UpdatePageView(id int) {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(update_pv)")
	}
	defer db.Close()

	var body Body
	db.Where("wiki_id = ?", id).First(&body)
	body.PV += 1
	db.Debug().Save(&body)
}
