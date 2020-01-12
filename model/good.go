package model

import "github.com/jinzhu/gorm"

/*
 * Goodする
 */
func CreateGood(user string, id int) {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(CreateGood())")
	}
	defer db.Close()

	var good Good
	good.GoodUser = user
	good.WikiID = id
	db.Debug().Create(&good)

}

func GetAllGood() []Good {
	db, err := gorm.Open("mysql", "gorm:password@/flashcard?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic("failed to connect database(CreateGood())")
	}
	defer db.Close()

	var good []Good
	db.Debug().Find(&good)

	return good
}
