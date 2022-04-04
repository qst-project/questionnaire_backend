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
	Order           string          `gorm:"not null;size:256"`
	Kind            string          `gorm:"not null;size:2"`
	Text            string          `gorm:"size:256;not null"`
}

type RadioPossibleAnswerDB struct {
	gorm.Model

	QuestionId uint
	Question   QuestionDB `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text       string     `gorm:"size:256;not null"`
}

type CheckboxPossibleAnswerDB struct {
	gorm.Model

	QuestionId uint
	Question   QuestionDB `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text       string     `gorm:"size:256;not null"`
}

type TextPossibleAnswerDB struct {
	gorm.Model

	QuestionId  uint
	Question    QuestionDB `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text        string     `gorm:"size:256;not null"`
	Placeholder string     `gorm:"size:256;"`
}

type RadioAnswerDB struct {
	gorm.Model

	QuestionnaireId       uint
	Questionnaire         QuestionnaireDB `gorm:"foreignKey:QuestionnaireId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	QuestionId            uint
	Question              QuestionDB `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RadioPossibleAnswerId uint
	RadioPossibleAnswer   RadioPossibleAnswerDB `gorm:"foreignKey:RadioPossibleAnswerId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CheckboxAnswerDB struct {
	gorm.Model

	QuestionnaireId          uint
	Questionnaire            QuestionnaireDB `gorm:"foreignKey:QuestionnaireId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	QuestionId               uint
	Question                 QuestionDB `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CheckboxPossibleAnswerId uint
	CheckboxPossibleAnswer   CheckboxPossibleAnswerDB `gorm:"foreignKey:CheckboxPossibleAnswerId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TextAnswerDB struct {
	gorm.Model

	QuestionnaireId      uint
	Questionnaire        QuestionnaireDB `gorm:"foreignKey:QuestionnaireId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	QuestionId           uint
	Question             QuestionDB `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TextPossibleAnswerId uint
	TextPossibleAnswer   TextPossibleAnswerDB `gorm:"foreignKey:TextPossibleAnswerId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Answer               string               `gorm:"size:256;not null"`
}
