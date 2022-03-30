package models

import (
	"github.com/qst-project/backend.git/pkg/api"
	"log"
	"strconv"
)

func (em *Questionnaire) GetgRPCModel() api.Survey {
	return api.Survey{
		Id:    strconv.FormatUint(uint64(em.ID), 10),
		Title: em.Title,
		Ref:   em.Ref,
	}
}

func (em *Questionnaire) From(gRRCModel *api.Survey) {
	id, err := strconv.ParseUint(gRRCModel.GetId(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.ID = uint(id)
	em.Title = gRRCModel.Title
	em.Ref = gRRCModel.Ref
}

func (em *Question) GetgRPCModel() api.Question {
	return api.Question{
		Id:              strconv.FormatUint(uint64(em.ID), 10),
		QuestionnaireId: strconv.FormatUint(uint64(em.QuestionnaireId), 10),
		Text:            em.Text,
		Order:           em.Order,
		Kind:            em.Kind,
	}
}

func (em *Question) From(gRRCModel *api.Question) {
	id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	QuestionnaireId, err2 := strconv.ParseUint(gRRCModel.QuestionnaireId, 10, 64)
	if err2 != nil {
		log.Fatal(err2)
	}
	em.ID = uint(id)
	em.QuestionnaireId = uint(QuestionnaireId)
	em.Text = gRRCModel.Text
	em.Order = gRRCModel.Order
	em.Kind = gRRCModel.Kind
}

// ------------------------------------------------------------------------

func (em *RadioPossibleAnswer) GetgRPCModel() api.RadioPossibleAnswer {
	return api.RadioPossibleAnswer{
		Id:         strconv.FormatUint(uint64(em.ID), 10),
		QuestionId: strconv.FormatUint(uint64(em.QuestionId), 10),
		Text:       em.Text,
	}
}

func (em *RadioPossibleAnswer) From(gRRCModel *api.RadioPossibleAnswer) {
	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	QuestionId, err := strconv.ParseUint(gRRCModel.QuestionId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	em.ID = uint(Id)
	em.QuestionId = uint(QuestionId)
	em.Text = gRRCModel.Text
}

func (em *CheckboxPossibleAnswer) GetgRPCModel() api.CheckboxPossibleAnswer {
	return api.CheckboxPossibleAnswer{
		Id:         strconv.FormatUint(uint64(em.ID), 10),
		QuestionId: strconv.FormatUint(uint64(em.QuestionId), 10),
		Text:       em.Text,
	}
}

func (em *CheckboxPossibleAnswer) From(gRRCModel *api.CheckboxPossibleAnswer) {
	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	QuestionId, err := strconv.ParseUint(gRRCModel.QuestionId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.ID = uint(Id)
	em.QuestionId = uint(QuestionId)
	em.Text = gRRCModel.Text
}

func (em *TextPossibleAnswer) GetgRPCModel() api.TextPossibleAnswer {
	return api.TextPossibleAnswer{
		Id:          strconv.FormatUint(uint64(em.ID), 10),
		QuestionId:  strconv.FormatUint(uint64(em.QuestionId), 10),
		Text:        em.Text,
		Placeholder: em.Placeholder,
	}
}

func (em *TextPossibleAnswer) From(gRRCModel *api.TextPossibleAnswer) {
	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	QuestionId, err := strconv.ParseUint(gRRCModel.QuestionId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.ID = uint(Id)
	em.QuestionId = uint(QuestionId)
	em.Text = gRRCModel.Text
	em.Placeholder = gRRCModel.Placeholder
}

// --------------------------------------------------------------------------

func (em *RadioAnswer) GetgRPCModel() api.RadioAnswer {
	return api.RadioAnswer{
		Id:                    strconv.FormatUint(uint64(em.ID), 10),
		RadioPossibleAnswerId: strconv.FormatUint(uint64(em.RadioPossibleAnswerId), 10),
	}
}

func (em *RadioAnswer) From(gRRCModel *api.RadioAnswer) {
	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	RadioPossibleAnswerId, err := strconv.ParseUint(gRRCModel.RadioPossibleAnswerId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.ID = uint(Id)
	em.RadioPossibleAnswerId = uint(RadioPossibleAnswerId)
}

func (em *CheckboxAnswer) GetgRPCModel() api.CheckboxAnswer {
	return api.CheckboxAnswer{
		Id:                       strconv.FormatUint(uint64(em.ID), 10),
		CheckboxPossibleAnswerId: strconv.FormatUint(uint64(em.CheckboxPossibleAnswerId), 10),
	}
}

func (em *CheckboxAnswer) From(gRRCModel *api.CheckboxAnswer) {
	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	CheckboxPossibleAnswerId, err := strconv.ParseUint(gRRCModel.CheckboxPossibleAnswerId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.ID = uint(Id)
	em.CheckboxPossibleAnswerId = uint(CheckboxPossibleAnswerId)
}

func (em *TextAnswer) GetgRPCModel() api.TextAnswer {
	return api.TextAnswer{
		Id:                   strconv.FormatUint(uint64(em.ID), 10),
		TextPossibleAnswerId: strconv.FormatUint(uint64(em.TextPossibleAnswerId), 10),
	}
}

func (em *TextAnswer) From(gRRCModel *api.TextAnswer) {
	Id, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	TextPossibleAnswerId, err := strconv.ParseUint(gRRCModel.TextPossibleAnswerId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	em.ID = uint(Id)
	em.TextPossibleAnswerId = uint(TextPossibleAnswerId)
}

func (em *UtilitySurvey) From(gRRCModel *api.Survey) {

	em.Questionnaire.From(gRRCModel)

	for _, question := range gRRCModel.GetQuestions() {
		var modelsQuestion Question
		modelsQuestion.From(question)
		em.Questions = append(em.Questions, &modelsQuestion)
	}

	for _, radioPossibleAnswer := range gRRCModel.GetRadioPossibleAnswers() {
		var modelRadioPossibleAnswer RadioPossibleAnswer
		modelRadioPossibleAnswer.From(radioPossibleAnswer)
		em.RadioPossibleAnswers = append(em.RadioPossibleAnswers, &modelRadioPossibleAnswer)
	}
	for _, checkboxPossibleAnswer := range gRRCModel.GetCheckboxPossibleAnswers() {
		var modelCheckboxPossibleAnswer CheckboxPossibleAnswer
		modelCheckboxPossibleAnswer.From(checkboxPossibleAnswer)
		em.CheckboxPossibleAnswers = append(em.CheckboxPossibleAnswers, &modelCheckboxPossibleAnswer)
	}
	for _, textPossibleAnswer := range gRRCModel.GetTextPossibleAnswers() {
		var modelTextPossibleAnswer TextPossibleAnswer
		modelTextPossibleAnswer.From(textPossibleAnswer)
		em.TextPossibleAnswers = append(em.TextPossibleAnswers, &modelTextPossibleAnswer)
	}

	for _, radioAnswer := range gRRCModel.GetRadioAnswers() {
		var modelRadioAnswer RadioAnswer
		modelRadioAnswer.From(radioAnswer)
		em.RadioAnswers = append(em.RadioAnswers, &modelRadioAnswer)
	}
	for _, checkboxAnswer := range gRRCModel.GetCheckboxAnswers() {
		var modelCheckboxAnswer CheckboxAnswer
		modelCheckboxAnswer.From(checkboxAnswer)
		em.CheckboxAnswers = append(em.CheckboxAnswers, &modelCheckboxAnswer)
	}
	for _, textAnswer := range gRRCModel.GetTextAnswers() {
		var modelTextAnswer TextAnswer
		modelTextAnswer.From(textAnswer)
		em.TextAnswers = append(em.TextAnswers, &modelTextAnswer)
	}
}

func (em *UtilitySurvey) GetgRPCModel() api.Survey {
	var apiSurvey api.Survey

	apiSurvey.Id = strconv.FormatUint(uint64(em.Questionnaire.ID), 10)
	apiSurvey.Title = em.Questionnaire.Title
	apiSurvey.Ref = em.Questionnaire.Ref

	for _, modelQuestion := range em.Questions {
		apiQuestion := modelQuestion.GetgRPCModel()
		apiSurvey.Questions = append(apiSurvey.Questions, &apiQuestion)
	}

	for _, modelRadioPossibleAnswer := range em.RadioPossibleAnswers {
		apiRadioPossibleAnswer := modelRadioPossibleAnswer.GetgRPCModel()
		apiSurvey.RadioPossibleAnswers = append(apiSurvey.RadioPossibleAnswers, &apiRadioPossibleAnswer)
	}
	for _, modelCheckboxPossibleAnswer := range em.CheckboxPossibleAnswers {
		apiCheckboxPossibleAnswer := modelCheckboxPossibleAnswer.GetgRPCModel()
		apiSurvey.CheckboxPossibleAnswers = append(apiSurvey.CheckboxPossibleAnswers, &apiCheckboxPossibleAnswer)
	}
	for _, modelTextPossibleAnswer := range em.TextPossibleAnswers {
		apiTextPossibleAnswer := modelTextPossibleAnswer.GetgRPCModel()
		apiSurvey.TextPossibleAnswers = append(apiSurvey.TextPossibleAnswers, &apiTextPossibleAnswer)
	}

	for _, modelRadioAnswer := range em.RadioAnswers {
		apiRadioAnswer := modelRadioAnswer.GetgRPCModel()
		apiSurvey.RadioAnswers = append(apiSurvey.RadioAnswers, &apiRadioAnswer)
	}
	for _, modelCheckboxAnswer := range em.CheckboxAnswers {
		apiCheckboxAnswer := modelCheckboxAnswer.GetgRPCModel()
		apiSurvey.CheckboxAnswers = append(apiSurvey.CheckboxAnswers, &apiCheckboxAnswer)
	}
	for _, modelTextAnswer := range em.TextAnswers {
		apiTextAnswer := modelTextAnswer.GetgRPCModel()
		apiSurvey.TextAnswers = append(apiSurvey.TextAnswers, &apiTextAnswer)
	}

	return apiSurvey
}
