package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var DBName = "mysql"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "gorm",
		DBName:   "flashcard",
		Password: "password",
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	) + "?charset=utf8&parseTime=True&loc=Asia%2FTokyo"
}

func GetUsingDBName() string {
	return DBName
}
