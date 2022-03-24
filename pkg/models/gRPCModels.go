package models

import (
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/models"
	"log"
	"strconv"
)

func (em *models.Questionnaire) GetgRPCModel() api.Survey {
	return api.Survey{
		Id: strconv.FormatUint(uint64(em.ID), 10),
		Title: em.Title,
		Ref: em.Ref,
		// Questions: nil,
		// RadioPossibleAnswers: nil,
		// CheckboxPossibleAnswers: nil,
		// TextPossibleAnswers: nil,
		// RadioAnswers: nil,
		// CheckboxAnswers: nil,
		// TextAnswers: nil,
		// get Question
		// get answer
	}
}

func (em *models.Questionnaire) From(gRRCModel api.Survey) {
	id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.ID = uint(id)
	em.Title = gRRCModel.Title
	em.Ref = gRPCModel.Ref
}

func (em *models.Question) GetgRPCModel() api.Question {
	return api.Question{
		Id: strconv.FormatUint(uint64(em.ID), 10),
		QuestionnaireId: strconv.FormatUint(uint64(em.QuestionnaireId), 10),
		Text: em.Title,
		Order: em.Order,
		Kind: em.Kind,
	}
}

func (em *models.Question) From(gRRCModel api.Question) {
	ID, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	QuestionnaireId, err2 := strconv.ParseUint(gRRCModel.QuestionnaireId, 10, 64)
	if err2 != nil {
		log.Fatal(err2)
	}
	em.ID = uint(ID)
	em.QuestionnaireId = uint(QuestionnaireId)
	em.Text = gRRCModel.Text
	em.Order = gRRCModel.Order
	em.Kind = gRRCModel.Kind
}

// ------------------------------------------------------------------------

func (em *models.RadioPossibleAnswer) GetgRPCModel() api.RadioPossibleAnswer {
	return api.RadioPossibleAnswer{
		Id: strconv.FormatUint(uint64(em.ID), 10),
		QuestionId: strconv.FormatUint(uint64(em.QuestionId), 10),
		Text: em.Title,
	}
}

func (em *models.RadioPossibleAnswer) From(gRRCModel api.RadioPossibleAnswer) {
	QuestionId, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.QuestionId = uint(QuestionId)
	em.Text = gRPCModel.Text
}

func (em *models.CheckboxPossibleAnswer) GetgRPCModel() api.CheckboxPossibleAnswer {
	return api.CheckboxPossibleAnswer{
		Id: strconv.FormatUint(uint64(em.ID), 10),
		QuestionId: strconv.FormatUint(uint64(em.QuestionId), 10),
		Text: em.Title,
	}
}

func (em *models.CheckboxPossibleAnswer) From(gRRCModel api.CheckboxPossibleAnswer) {
	QuestionId, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.QuestionId = uint(QuestionId)
	em.Text = gRPCModel.Text
}

func (em *models.TextPossibleAnswer) GetgRPCModel() api.TextPossibleAnswer {
	return api.TextPossibleAnswer{
		Id: strconv.FormatUint(uint64(em.ID), 10),
		QuestionId: strconv.FormatUint(uint64(em.QuestionId), 10),
		Text: em.Title,
	}
}

func (em *models.TextPossibleAnswer) From(gRRCModel api.TextPossibleAnswer) {
	QuestionId, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.QuestionId = uint(QuestionId)
	em.Text = gRPCModel.Text
	em.Placeholder = gRPCModel.Placeholder
}

// --------------------------------------------------------------------------

func (em *models.RadioAnswer) GetgRPCModel() api.RadioAnswer {
	return api.RadioAnswer{
		Id: strconv.FormatUint(uint64(em.ID), 10),
		RadioPossibleAnswerId: strconv.FormatUint(uint64(em.RadioPossibleAnswerId), 10),
	}
}

func (em *models.RadioAnswer) From(gRRCModel api.RadioAnswer) {
	RadioPossibleAnswerId, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.RadioPossibleAnswerId = uint(RadioPossibleAnswerId)
}

func (em *models.CheckboxAnswer) GetgRPCModel() api.CheckboxAnswer {
	return api.CheckboxAnswer{
		Id: strconv.FormatUint(uint64(em.ID), 10),
		CheckboxPossibleAnswerId: strconv.FormatUint(uint64(em.CheckboxPossibleAnswerId), 10),
	}
}

func (em *models.CheckboxAnswer) From(gRRCModel api.CheckboxAnswer) {
	CheckboxPossibleAnswerId, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.CheckboxPossibleAnswerId = uint(CheckboxPossibleAnswerId)
}

func (em *models.TextAnswer) GetgRPCModel() api.TextAnswer {
	return api.TextAnswer{
		Id: strconv.FormatUint(uint64(em.ID), 10),
		TextPossibleAnswerId: strconv.FormatUint(uint64(em.TextPossibleAnswerId), 10),
	}
}

func (em *models.TextAnswer) From(gRRCModel api.TextAnswer) {
	TextPossibleAnswerId, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.TextPossibleAnswerId = uint(TextPossibleAnswerId)
}

