package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	databse, err := gorm.Open(mysql.Open("root:Ddch@ntiqu3 @tcp(localhost:3306)/go_ruang_bengkel"))
	if err != nil {
		panic(err)
	}

	databse.AutoMigrate(&User{})

	DB = databse
}
