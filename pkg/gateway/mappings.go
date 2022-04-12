package gateway

import (
	"github.com/qst-project/backend.git/pkg/core"
)

func (em *QuestionnaireDB) ToCore() core.Questionnaire {
	return core.Questionnaire{
		Id:    core.Id(em.ID),
		Title: em.Title,
		Ref:   em.Ref,
	}
}

func (em QuestionnaireDB) FromCore(questionnaire *core.Questionnaire) {
	em.ID = uint(questionnaire.Id)
	em.Title = questionnaire.Title
	em.Ref = questionnaire.Ref
}

func (em *QuestionDB) ToCore() core.Question {
	return core.Question{
		Id:    core.Id(em.ID),
		Text:  em.Text,
		Order: em.Order,
		Kind:  em.Kind,
	}
}

func (em *QuestionDB) FromCore(question *core.Question) {
	em.ID = uint(question.Id)
	em.Text = question.Text
	em.Order = question.Order
	em.Kind = question.Kind
}

// ------------------------------------------------------------------------

func (em *RadioPossibleAnswerDB) ToCore() core.RadioPossibleAnswer {
	return core.RadioPossibleAnswer{
		Id:         core.Id(em.ID),
		QuestionId: em.QuestionId,
		Text:       em.Text,
	}
}

func (em *RadioPossibleAnswerDB) FromCore(radioPossibleAnswer *core.RadioPossibleAnswer) {
	em.ID = uint(radioPossibleAnswer.Id)
	em.QuestionId = radioPossibleAnswer.QuestionId
	em.Text = radioPossibleAnswer.Text
}

func (em *CheckboxPossibleAnswerDB) ToCore() core.CheckboxPossibleAnswer {
	return core.CheckboxPossibleAnswer{
		Id:         core.Id(em.ID),
		QuestionId: em.QuestionId,
		Text:       em.Text,
	}
}

func (em *CheckboxPossibleAnswerDB) FromCore(checkboxPossibleAnswer *core.CheckboxPossibleAnswer) {
	em.ID = uint(checkboxPossibleAnswer.Id)
	em.QuestionId = checkboxPossibleAnswer.QuestionId
	em.Text = checkboxPossibleAnswer.Text
}

func (em *TextPossibleAnswerDB) ToCore() core.TextPossibleAnswer {
	return core.TextPossibleAnswer{
		Id:          core.Id(em.ID),
		QuestionId:  em.QuestionId,
		Text:        em.Text,
		Placeholder: em.Placeholder,
	}
}

func (em *TextPossibleAnswerDB) FromCore(textPossibleAnswer *core.TextPossibleAnswer) {
	em.ID = uint(textPossibleAnswer.Id)
	em.QuestionId = textPossibleAnswer.QuestionId
	em.Text = textPossibleAnswer.Text
	em.Placeholder = textPossibleAnswer.Placeholder
}

// --------------------------------------------------------------------------

//func (em *RadioAnswerDB) ToCore() core.RadioAnswer {
//	return core.RadioAnswer{
//		Id:                    core.Id(em.ID),
//		QuestionId:            em.QuestionId,
//		QuestionnaireId:       em.QuestionnaireId,
//		RadioPossibleAnswerId: em.RadioPossibleAnswerId,
//	}
//}
//
//func (em *RadioAnswerDB) FromCore(radioAnswer *core.RadioAnswer) {
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
