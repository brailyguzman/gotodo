package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
}

func (User) Create(db *gorm.DB, user *User) error {
	var err error
	return err
}
