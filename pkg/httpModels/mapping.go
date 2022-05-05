package httpModels

import (
	"github.com/qst-project/backend.git/pkg/core"
)

func (em *Questionnaire) ToCore() core.Questionnaire {
	var coreQuestions []*core.Question

	for _, protoQuestion := range em.Questions {
		coreQuestion := protoQuestion.ToCore()
		coreQuestions = append(coreQuestions, &coreQuestion)
	}
	return core.Questionnaire{
		Id:        core.Id(em.Id),
		Title:     em.Title,
		Ref:       em.Ref,
		Questions: coreQuestions,
	}
}

func (em *Questionnaire) FromCore(questionnaire *core.Questionnaire) {
	em.Id = uint(questionnaire.Id)
	em.Title = questionnaire.Title
	em.Ref = questionnaire.Ref
	for _, coreQuestion := range questionnaire.Questions {
		var protoQuestion Question
		protoQuestion.FromCore(coreQuestion)
		em.Questions = append(em.Questions, &protoQuestion)
	}
}

func (em *Question) ToCore() core.Question {
	return core.Question{
		Id:        core.Id(em.Id),
		Statement: em.Statement,
		Type:      core.Type(em.Type),
		Options:   em.Options,
	}
}

func (em *Question) FromCore(question *core.Question) {
	em.Id = uint(question.Id)
	em.Statement = question.Statement
	em.Type = Type(question.Type)
	em.Options = question.Options
}
