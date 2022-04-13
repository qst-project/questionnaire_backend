package gateway

import (
	"github.com/qst-project/backend.git/pkg/core"
)

func (em *QuestionnaireDB) ToCore(questions []*QuestionDB, options []*OptionsDB) core.Questionnaire {
	var coreQuestions []*core.Question
	for _, questionDB := range questions {
		coreQuestions = append(coreQuestions, questionDB.ToCore(options))
	}
	return core.Questionnaire{
		Id:        core.Id(em.ID),
		Title:     em.Title,
		Ref:       em.Ref,
		Questions: coreQuestions,
	}
}

func (em *QuestionnaireDB) FromCore(questionnaire *core.Questionnaire) {
	em.ID = uint(questionnaire.Id)
	em.Title = questionnaire.Title
	em.Ref = questionnaire.Ref
}

func (em *QuestionDB) ToCore(options []*OptionsDB) *core.Question {
	var coreOptions []string
	for _, optionDB := range options {
		coreOptions = append(coreOptions, optionDB.ToCore())
	}
	return &core.Question{
		Id:        core.Id(em.ID),
		Statement: em.Statement,
		Type:      em.Type,
		Options:   coreOptions,
	}
}

func (em *QuestionDB) FromCore(order, questionnaireId uint, question *core.Question) {
	em.ID = uint(question.Id)
	em.QuestionnaireId = questionnaireId
	em.Type = question.Type
	em.Statement = question.Statement
	em.Order = order
}

func (em *OptionsDB) ToCore() string {
	return em.Text
}

func (em *OptionsDB) FromCore(option string, questionId uint) {
	em.QuestionId = questionId
	em.Text = option
}
