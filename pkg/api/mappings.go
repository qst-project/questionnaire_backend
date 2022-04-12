package api

import (
	"github.com/qst-project/backend.git/pkg/core"
	"strconv"
)

func (em *Questionnaire) ToCore() (qst core.Questionnaire, err error) {
	qst = core.Questionnaire{
		Id:    core.Id(em.Id),
		Title: em.Title,
		Ref:   em.Ref,
	}
	//TODO ...
	return qst, nil
}

func (em Questionnaire) FromCore(questionnaire *core.Questionnaire) {
	em.Id = uint32(questionnaire.Id)
	em.Title = questionnaire.Title
	em.Ref = questionnaire.Ref

	//TODO ...
}

func (em *Question) ToCore(order string) core.Question {
	id, _ := strconv.ParseUint(em.Id, 10, 32)
	return core.Question{
		Id:    core.Id(id),
		Text:  em.Statement,
		Kind:  em.Type,
		Order: order,
	}
}

func (em *Question) FromCore(question *core.Question) {
	em.Id = strconv.FormatUint(uint64(question.Id), 32)
	em.QuestionnaireId = strconv.FormatUint(uint64(question.QuestionnaireId), 32)
	em.Statement = question.Text
	em.Type = question.Kind
	//TODO ...
}

// ------------------------------------------------------------------------

func (em *RadioPossibleAnswer) ToCore() core.RadioPossibleAnswer {
	id, _ := strconv.ParseUint(em.Id, 10, 32)
	return core.RadioPossibleAnswer{
		Id:         core.Id(id),
		QuestionId: em.QuestionId,
		Text:       em.Text,
	}
}

func (em *RadioPossibleAnswer) FromCore(radioPossibleAnswer *core.RadioPossibleAnswer) {
	em.ID = uint(radioPossibleAnswer.Id)
	em.QuestionId = radioPossibleAnswer.QuestionId
	em.Text = radioPossibleAnswer.Text
}

func (em *CheckboxPossibleAnswer) ToCore() core.CheckboxPossibleAnswer {
	return core.CheckboxPossibleAnswer{
		Id:         core.Id(em.ID),
		QuestionId: em.QuestionId,
		Text:       em.Text,
	}
}

func (em *CheckboxPossibleAnswer) FromCore(checkboxPossibleAnswer *core.CheckboxPossibleAnswer) {
	em.ID = uint(checkboxPossibleAnswer.Id)
	em.QuestionId = checkboxPossibleAnswer.QuestionId
	em.Text = checkboxPossibleAnswer.Text
}

func (em *TextPossibleAnswer) ToCore() core.TextPossibleAnswer {
	return core.TextPossibleAnswer{
		Id:          core.Id(em.ID),
		QuestionId:  em.QuestionId,
		Text:        em.Text,
		Placeholder: em.Placeholder,
	}
}

func (em *TextPossibleAnswer) FromCore(textPossibleAnswer *core.TextPossibleAnswer) {
	em.ID = uint(textPossibleAnswer.Id)
	em.QuestionId = textPossibleAnswer.QuestionId
	em.Text = textPossibleAnswer.Text
	em.Placeholder = textPossibleAnswer.Placeholder
}

// --------------------------------------------------------------------------
//
//func (em *RadioAnswer) ToCore() core.RadioAnswer {
//	return core.RadioAnswer{
//		Id:                    core.Id(em.ID),
//		QuestionId:            em.QuestionId,
//		QuestionnaireId:       em.QuestionnaireId,
//		RadioPossibleAnswerId: em.RadioPossibleAnswerId,
//	}
//}
//
//func (em *RadioAnswer) FromCore(radioAnswer *core.RadioAnswer) {
//	em.ID = uint(radioAnswer.Id)
//	em.QuestionId = radioAnswer.QuestionId
//	em.QuestionnaireId = radioAnswer.QuestionnaireId
//	em.RadioPossibleAnswerId = radioAnswer.RadioPossibleAnswerId
//}
//
//func (em *CheckboxAnswerDB) ToCore() core.CheckboxAnswer {
//	return core.CheckboxAnswer{
//		Id:                       core.Id(em.ID),
//		QuestionId:               em.QuestionId,
//		QuestionnaireId:          em.QuestionnaireId,
//		CheckboxPossibleAnswerId: em.CheckboxPossibleAnswerId,
//	}
//}
//
//func (em *CheckboxAnswerDB) From(checkboxAnswer *core.CheckboxAnswer) {
//	em.ID = uint(checkboxAnswer.Id)
//	em.CheckboxPossibleAnswerId = checkboxAnswer.CheckboxPossibleAnswerId
//	em.QuestionnaireId = checkboxAnswer.QuestionnaireId
//	em.QuestionId = checkboxAnswer.QuestionId
//}
//
//func (em *TextAnswerDB) GetGRPCModel() core.TextAnswer {
//	return core.TextAnswer{
//		Id:                   core.Id(em.ID),
//		QuestionId:           em.QuestionId,
//		QuestionnaireId:      em.QuestionnaireId,
//		TextPossibleAnswerId: em.TextPossibleAnswerId,
//		Answer:               em.Answer,
//	}
//}
//
//func (em *TextAnswerDB) From(textAnswer *core.TextAnswer) {
//	em.ID = uint(textAnswer.Id)
//	em.TextPossibleAnswerId = textAnswer.TextPossibleAnswerId
//	em.QuestionId = textAnswer.QuestionId
//	em.QuestionnaireId = textAnswer.QuestionnaireId
//	em.Answer = textAnswer.Answer
//}
