package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//d, err := gorm.Open("mysql", "ahmed:gomysql@123/bookstore?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "doadmin:AVNS_PED5LfbsoGZqR9x@tcp(bookstore-do-user-8417367-0.b.db.ondigitalocean.com:25060)/defaultdb")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
