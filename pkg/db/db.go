package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(conf *viper.Viper) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.dns")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
