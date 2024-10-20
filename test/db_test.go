package test

import (
	"github.com/rizky201008/subscription-tracker-api/models/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	host := Viper.GetString("db_test.host")
	port := Viper.GetString("db_test.port")
	user := Viper.GetString("db_test.user")
	password := Viper.GetString("db_test.password")
	database := Viper.GetString("db_test.database")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect databases" + err.Error())
	}
	err = db.AutoMigrate(&domain.Setting{}, &domain.Subscription{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

var Db = InitDb()
