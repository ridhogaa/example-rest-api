package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:100;not null" json:"name"`
	Email string `gorm:"size:100;unique;not null" json:"email"`
}
