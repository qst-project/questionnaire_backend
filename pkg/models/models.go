package models

import (
	"gorm.io/gorm"
)

type Questionnaire struct {
	gorm.Model
	Title string `gorm:"not null;size:256"`
	Ref   string `gorm:"not null;size:256"`
}

type Question struct {
	gorm.Model
	QuestionnaireId uint
	Questionnaire   Questionnaire `gorm:"foreignKey:QuestionnaireId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Order           string        `gorm:"not null;size:256"`
	Kind            string        `gorm:"not null;size:2"`
	Text            string        `gorm:"size:256;not null"`
}

type RadioPossibleAnswer struct {
	gorm.Model
	QuestionId uint
	Question   Question `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text       string   `gorm:"size:256;not null"`
}

type CheckboxPossibleAnswer struct {
	gorm.Model
	QuestionId uint
	Question   Question `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text       string   `gorm:"size:256;not null"`
}

type TextPossibleAnswer struct {
	gorm.Model
	QuestionId  uint
	Question    Question `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Text        string   `gorm:"size:256;not null"`
	Placeholder string   `gorm:"size:256;"`
}

type RadioAnswer struct {
	gorm.Model
	QuestionnaireId       uint
	Questionnaire         Questionnaire `gorm:"foreignKey:QuestionnaireId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	QuestionId            uint
	Question              Question `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RadioPossibleAnswerId uint
	RadioPossibleAnswer   RadioPossibleAnswer `gorm:"foreignKey:RadioPossibleAnswerId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CheckboxAnswer struct {
	gorm.Model
	QuestionnaireId          uint
	Questionnaire            Questionnaire `gorm:"foreignKey:QuestionnaireId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	QuestionId               uint
	Question                 Question `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CheckboxPossibleAnswerId uint
	CheckboxPossibleAnswer   CheckboxPossibleAnswer `gorm:"foreignKey:CheckboxPossibleAnswerId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TextAnswer struct {
	gorm.Model
	QuestionnaireId      uint
	Questionnaire        Questionnaire `gorm:"foreignKey:QuestionnaireId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	QuestionId           uint
	Question             Question `gorm:"foreignKey:QuestionId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TextPossibleAnswerId uint
	TextPossibleAnswer   TextPossibleAnswer `gorm:"foreignKey:TextPossibleAnswerId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Answer               string             `gorm:"size:256;not null"`
}

type UtilitySurvey struct {
	Questionnaire           Questionnaire
	Questions               []*Question
	RadioPossibleAnswers    []*RadioPossibleAnswer
	CheckboxPossibleAnswers []*CheckboxPossibleAnswer
	TextPossibleAnswers     []*TextPossibleAnswer
	//RadioAnswers            []*RadioAnswer
	//CheckboxAnswers         []*CheckboxAnswer
	//TextAnswers             []*TextAnswer
}
