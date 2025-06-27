package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Text   string `json:"text" gorm:"not null"`
	Done   bool   `json:"done" gorm:"default:false"`
	UserID uint   `json:"user_id" gorm:"not null"`
}
