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

//func (em *QuestionDB) GetgRPCModel() api.Question {
//	return api.Question{
//		Id:              strconv.FormatUint(uint64(em.ID), 10),
//		QuestionnaireId: strconv.FormatUint(uint64(em.QuestionnaireId), 10),
//		Text:            em.Text,
//		Order:           em.Order,
//		Kind:            em.Kind,
//	}
//}
//
//func (em *QuestionDB) From(gRRCModel *api.Question) {
//	id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	QuestionnaireId, err2 := strconv.ParseUint(gRRCModel.QuestionnaireId, 10, 64)
//	if err2 != nil {
//		log.Fatal(err2)
//	}
//	em.ID = uint(id)
//	em.QuestionnaireId = uint(QuestionnaireId)
//	em.Text = gRRCModel.Text
//	em.Order = gRRCModel.Order
//	em.Kind = gRRCModel.Kind
//}

// ------------------------------------------------------------------------

//func (em *RadioPossibleAnswerDB) GetgRPCModel() api.RadioPossibleAnswer {
//	return api.RadioPossibleAnswer{
//		Id:         strconv.FormatUint(uint64(em.ID), 10),
//		QuestionId: strconv.FormatUint(uint64(em.QuestionId), 10),
//		Text:       em.Text,
//	}
//}
//
//func (em *RadioPossibleAnswerDB) From(gRRCModel *api.RadioPossibleAnswer) {
//	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	QuestionId, err := strconv.ParseUint(gRRCModel.QuestionId, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	em.ID = uint(Id)
//	em.QuestionId = uint(QuestionId)
//	em.Text = gRRCModel.Text
//}
//
//func (em *CheckboxPossibleAnswerDB) GetgRPCModel() api.CheckboxPossibleAnswer {
//	return api.CheckboxPossibleAnswer{
//		Id:         strconv.FormatUint(uint64(em.ID), 10),
//		QuestionId: strconv.FormatUint(uint64(em.QuestionId), 10),
//		Text:       em.Text,
//	}
//}
//
//func (em *CheckboxPossibleAnswerDB) From(gRRCModel *api.CheckboxPossibleAnswer) {
//	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	QuestionId, err := strconv.ParseUint(gRRCModel.QuestionId, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	em.ID = uint(Id)
//	em.QuestionId = uint(QuestionId)
//	em.Text = gRRCModel.Text
//}
//
//func (em *TextPossibleAnswerDB) GetgRPCModel() api.TextPossibleAnswer {
//	return api.TextPossibleAnswer{
//		Id:          strconv.FormatUint(uint64(em.ID), 10),
//		QuestionId:  strconv.FormatUint(uint64(em.QuestionId), 10),
//		Text:        em.Text,
//		Placeholder: em.Placeholder,
//	}
//}
//
//func (em *TextPossibleAnswerDB) From(gRRCModel *api.TextPossibleAnswer) {
//	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	QuestionId, err := strconv.ParseUint(gRRCModel.QuestionId, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	em.ID = uint(Id)
//	em.QuestionId = uint(QuestionId)
//	em.Text = gRRCModel.Text
//	em.Placeholder = gRRCModel.Placeholder
//}

// --------------------------------------------------------------------------

//func (em *RadioAnswerDB) GetgRPCModel() api.RadioAnswer {
//	return api.RadioAnswer{
//		Id:                  strconv.FormatUint(uint64(em.ID), 10),
//		QuestionId:          strconv.FormatUint(uint64(em.QuestionId), 10),
//		QuestionnaireId:     strconv.FormatUint(uint64(em.QuestionnaireId), 10),
//		RadioPossibleAnswer: strconv.FormatUint(uint64(em.RadioPossibleAnswerId), 10),
//	}
//}
//
//func (em *RadioAnswerDB) From(gRRCModel *api.RadioAnswer) {
//	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	RadioPossibleAnswerId, err := strconv.ParseUint(gRRCModel.RadioPossibleAnswer, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	em.ID = uint(Id)
//	em.RadioPossibleAnswerId = uint(RadioPossibleAnswerId)
//}
//
//func (em *CheckboxAnswerDB) GetgRPCModel() api.CheckboxAnswer {
//	return api.CheckboxAnswer{
//		Id:                     strconv.FormatUint(uint64(em.ID), 10),
//		QuestionId:             strconv.FormatUint(uint64(em.QuestionId), 10),
//		QuestionnaireId:        strconv.FormatUint(uint64(em.QuestionnaireId), 10),
//		CheckboxPossibleAnswer: strconv.FormatUint(uint64(em.CheckboxPossibleAnswerId), 10),
//	}
//}
//
//func (em *CheckboxAnswerDB) From(gRRCModel *api.CheckboxAnswer) {
//	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	CheckboxPossibleAnswerId, err := strconv.ParseUint(gRRCModel.CheckboxPossibleAnswer, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	em.ID = uint(Id)
//	em.CheckboxPossibleAnswerId = uint(CheckboxPossibleAnswerId)
//}
//
//func (em *TextAnswerDB) GetGRPCModel() api.TextAnswer {
//	return api.TextAnswer{
//		Id:                 strconv.FormatUint(uint64(em.ID), 10),
//		QuestionId:         strconv.FormatUint(uint64(em.QuestionId), 10),
//		QuestionnaireId:    strconv.FormatUint(uint64(em.QuestionnaireId), 10),
//		TextPossibleAnswer: strconv.FormatUint(uint64(em.TextPossibleAnswerId), 10),
//	}
//}
//
//func (em *TextAnswerDB) From(gRRCModel *api.TextAnswer) {
//	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	TextPossibleAnswerId, err := strconv.ParseUint(gRRCModel.TextPossibleAnswer, 10, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	em.ID = uint(Id)
//	em.TextPossibleAnswerId = uint(TextPossibleAnswerId)
//}
