package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id       int    `json:"ID" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type task task {
// 	gorm.Model
// 	Id			int 		`json:"ID" gorm:"primary_key"`
// 	Title		string		`json:"title"`
// 	Description string 		`json:"description"`
// 	DueDate		time.Time	`json:"due_date"`
// 	Status		string		`json:"status"`

// }
