package model

import (
	"flashcardsProject/config"

	"github.com/jinzhu/gorm"
)

/*
 * PV数の更新(+1)
 */
func UpdatePageView(id int) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))

	if err != nil {
		panic("failed to connect database(update_pv)")
	}

	var body Body
	config.DB.Where("wiki_id = ?", id).First(&body)
	body.PV += 1
	config.DB.Debug().Save(&body)
}
