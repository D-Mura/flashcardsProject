package model

import (
	"flashcardsProject/config"

	"github.com/jinzhu/gorm"
)

/*
 * Goodする
 */
func CreateGood(user string, id int) {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(CreateGood())")
	}

	var good Good
	good.GoodUser = user
	good.WikiID = id
	config.DB.Debug().Create(&good)

}

func GetAllGood() []Good {
	var err error
	config.DB, err = gorm.Open(config.GetUsingDBName(), config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic("failed to connect database(CreateGood())")
	}

	var good []Good
	config.DB.Debug().Find(&good)

	return good
}
