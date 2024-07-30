package models

import "gorm.io/gorm"

// GORM model in this project
type Fact struct {
	gorm.Model
	Question	string	`json:"question" gorm:"text;not null;default:null`
	Answer		string	`json:"answer" gorm:"text;not null;default:null`
}