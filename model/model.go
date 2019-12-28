package model

import "github.com/jinzhu/gorm"

type Wiki struct {
	gorm.Model
	Body        Body
	Title       string
	PictureName string
	ScreenID    int
	Good        int `gorm:"default:0"`
}

type Body struct {
	gorm.Model
	Text   string
	Author string
	Url    string
	WikiID int `gorm:"index"`
}

type User struct {
	gorm.Model
	Name     string
	Password string
}

type SessionInfo struct {
	UserId interface{}
}
