package gateway

import "gorm.io/gorm"

type QuestionnaireDB struct {
	gorm.Model

	Title string `gorm:"not null;size:256"`
	Ref   string `gorm:"not null;size:256"`
}

type QuestionDB struct {
	gorm.Model

	QuestionnaireId uint
	Questionnaire   QuestionnaireDB `gorm:"foreignKey:QuestionnaireId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Type            uint
	Statement       string `gorm:"size:256;not null"`
	Order           uint
}

type OptionsDB struct {
	gorm.Model

	QuestionId uint
	Question   QuestionDB `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text       string     `gorm:"size:256;not null"`
}
