package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Done        bool   `json:"done" gorm:"default:false"`
	UserID      uint   `json:"user_id" gorm:"not null"`
	User        User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

func (Todo) Create(db *gorm.DB, todo *Todo) error {
	var err error

	// TODO: Implement the logic to create a new todo item

	return err
}
