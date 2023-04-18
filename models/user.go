package models

type User struct {
	id       int64  `gorm:"primaryKey" json : "id"`
	fullName string `gorm:"type:varchar(300)" json : "full_name"`
	email    string `gorm:"type:varchar(300)" json : "email"`
}
