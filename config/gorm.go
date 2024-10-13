package config

import (
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	host := Vipers.GetString("db.host")
	port := Vipers.GetString("db.port")
	user := Vipers.GetString("db.user")
	password := Vipers.GetString("db.password")
	database := Vipers.GetString("db.database")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect databases" + err.Error())
	}
	err = db.AutoMigrate(&domain.Setting{}, &domain.Subscription{})
	if err != nil {
		panic(err.Error())
	}
	Db = db
}
