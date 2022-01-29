package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	googleId string `gorm:"not null;unique"`
}

type Questionnaire struct {
	gorm.Model
	title string `gorm:"not null;size:256"`
	creator_id User
}

type Question struct {
	gorm.Model
	order int `gorm:"not null"`
	questionnaire Questionnaire
	kind int `gorm:"not null"`
	text string `gorm:"size:256;not null"`

}

type RadioPossibleAnswer struct {
	gorm.Model
	question Question
	text string `gorm:"size:256;not null"`
}

type RadioAnswer struct {
	gorm.Model
	radioPossibleAnswer RadioPossibleAnswer
}

type CheckboxPossibleAnswer struct {
	gorm.Model
	question Question
	text string `gorm:"size:256;not null"`
}

type CheckboxAnswer struct {
	gorm.Model
	checkboxPossibleAnswer CheckboxPossibleAnswer
}

type TextPossibleAnswer struct {
	gorm.Model
	question Question
	text string `gorm:"size:256;not null"`
	placeholder string `gorm:"size:256;"`
}

type TextAnswer struct {
	gorm.Model
	textPossibleAnswer TextPossibleAnswer
	answer string `gorm:"size:256;not null"`
}