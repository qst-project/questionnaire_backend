package api

import (
	"github.com/qst-project/backend.git/pkg/core"
	"strconv"
)

func (em *Questionnaire) ToCore() core.Questionnaire {
	var coreQuestions []*core.Question

	for _, protoQuestion := range em.GetQuestions() {
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
	em.Id = uint32(questionnaire.Id)
	em.Title = questionnaire.Title
	em.Ref = questionnaire.Ref
	for _, coreQuestion := range questionnaire.Questions {
		var protoQuestion Question
		protoQuestion.FromCore(coreQuestion)
		em.Questions = append(em.Questions, &protoQuestion)
	}
}

func (em *Question) ToCore() core.Question {
	id, _ := strconv.ParseUint(em.Id, 10, 32)
	return core.Question{
		Id:        core.Id(id),
		Statement: em.Statement,
		Type:      em.Type,
		Options:   em.Options,
	}
}

func (em *Question) FromCore(question *core.Question) {
	em.Id = strconv.FormatUint(uint64(question.Id), 32)
	em.Statement = question.Statement
	em.Type = question.Type
	em.Options = question.Options
}
