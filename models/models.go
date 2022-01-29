package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	GoogleId string `gorm:"not null;unique"`
}

type Questionnaire struct {
	gorm.Model
	Title     string `gorm:"not null;size:256"`
	CreatorId User
}

type Question struct {
	gorm.Model
	Order int `gorm:"not null"`
	Questionnaire Questionnaire
	Kind int `gorm:"not null"`
	Text string `gorm:"size:256;not null"`

}

type RadioPossibleAnswer struct {
	gorm.Model
	Question Question
	Text string `gorm:"size:256;not null"`
}

type RadioAnswer struct {
	gorm.Model
	RadioPossibleAnswer RadioPossibleAnswer
}

type CheckboxPossibleAnswer struct {
	gorm.Model
	Question Question
	Text string `gorm:"size:256;not null"`
}

type CheckboxAnswer struct {
	gorm.Model
	CheckboxPossibleAnswer CheckboxPossibleAnswer
}

type TextPossibleAnswer struct {
	gorm.Model
	Question Question
	Text string `gorm:"size:256;not null"`
	Placeholder string `gorm:"size:256;"`
}

type TextAnswer struct {
	gorm.Model
	TextPossibleAnswer TextPossibleAnswer
	Answer string `gorm:"size:256;not null"`
}