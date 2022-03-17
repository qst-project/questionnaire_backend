package models

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

// type User struct {
// 	gorm.Model
// 	GoogleId string `gorm:"not null;unique"`
// }

type Questionnaire struct {
	gorm.Model
	Title  string `gorm:"not null;size:256"`
	Questions pq.StringArray `gorm:"type:text[]"`
	// UserId string
	// User   User `gorm:"foreignKey:UserId;references:GoogleId"`
}

type Question struct {
	gorm.Model
	QuestionnaireId uint
	Questionnaire   Questionnaire `gorm:"foreignKey:QuestionnaireId;references:ID"`
	Order           int           `gorm:"not null"`
	Kind            int           `gorm:"not null"`
	Text            string        `gorm:"size:256;not null"`
}

type RadioPossibleAnswer struct {
	gorm.Model
	QuestionId uint
	Question   Question `gorm:"foreignKey:QuestionId;references:ID"`
	Text       string   `gorm:"size:256;not null"`
}

type RadioAnswer struct {
	gorm.Model
	RadioPossibleAnswerId uint
	RadioPossibleAnswer   RadioPossibleAnswer `gorm:"foreignKey:RadioPossibleAnswerId;references:ID"`
}

type CheckboxPossibleAnswer struct {
	gorm.Model
	QuestionId uint
	Question   Question `gorm:"foreignKey:QuestionId;references:ID"`
	Text       string   `gorm:"size:256;not null"`
}

type CheckboxAnswer struct {
	gorm.Model
	CheckboxPossibleAnswerId uint
	CheckboxPossibleAnswer   CheckboxPossibleAnswer `gorm:"foreignKey:CheckboxPossibleAnswerId;references:ID"`
}

type TextPossibleAnswer struct {
	gorm.Model
	QuestionId  uint
	Question    Question `gorm:"foreignKey:QuestionId;references:ID"`
	Text        string   `gorm:"size:256;not null"`
	Placeholder string   `gorm:"size:256;"`
}

type TextAnswer struct {
	gorm.Model
	TextPossibleAnswerId uint
	TextPossibleAnswer   TextPossibleAnswer `gorm:"foreignKey:TextPossibleAnswerId;references:ID"`
	Answer               string             `gorm:"size:256;not null"`
}
