package unit

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/models"
	"testing"
)

func Add(a int, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

type UtilitySurvey struct {
	Questionnaire           *models.Questionnaire
	Questions               []*models.Question
	RadioPossibleAnswers    []*models.RadioPossibleAnswer
	CheckboxPossibleAnswers []*models.CheckboxPossibleAnswer
	TextPossibleAnswers     []*models.TextPossibleAnswer
	RadioAnswers            []*models.RadioAnswer
	CheckboxAnswers         []*models.CheckboxAnswer
	TextAnswers             []*models.TextAnswer
}

func (em *UtilitySurvey) From(gRRCModel *api.Survey) {
	em.Questionnaire.From(gRRCModel)

	for _, question := range gRRCModel.GetQuestions() {
		var modelsQuestion models.Question
		modelsQuestion.From(question)
		em.Questions = append(em.Questions, &modelsQuestion)
	}

	for _, radioPossibleAnswer := range gRRCModel.GetRadioPossibleAnswers() {
		var modelRadioPossibleAnswer models.RadioPossibleAnswer
		modelRadioPossibleAnswer.From(radioPossibleAnswer)
		em.RadioPossibleAnswers = append(em.RadioPossibleAnswers, &modelRadioPossibleAnswer)
	}
	for _, checkboxPossibleAnswer := range gRRCModel.GetCheckboxPossibleAnswers() {
		var modelCheckboxPossibleAnswer models.CheckboxPossibleAnswer
		modelCheckboxPossibleAnswer.From(checkboxPossibleAnswer)
		em.CheckboxPossibleAnswers = append(em.CheckboxPossibleAnswers, &modelCheckboxPossibleAnswer)
	}
	for _, textPossibleAnswer := range gRRCModel.GetTextPossibleAnswers() {
		var modelTextPossibleAnswer models.TextPossibleAnswer
		modelTextPossibleAnswer.From(textPossibleAnswer)
		em.TextPossibleAnswers = append(em.TextPossibleAnswers, &modelTextPossibleAnswer)
	}

	for _, radioAnswer := range gRRCModel.GetRadioAnswers() {
		var modelRadioAnswer models.RadioAnswer
		modelRadioAnswer.From(radioAnswer)
		em.RadioAnswers = append(em.RadioAnswers, &modelRadioAnswer)
	}
	for _, checkboxAnswer := range gRRCModel.GetCheckboxAnswers() {
		var modelCheckboxAnswer models.CheckboxAnswer
		modelCheckboxAnswer.From(checkboxAnswer)
		em.CheckboxAnswers = append(em.CheckboxAnswers, &modelCheckboxAnswer)
	}
	for _, textAnswer := range gRRCModel.GetTextAnswers() {
		var modelTextAnswer models.TextAnswer
		modelTextAnswer.From(textAnswer)
		em.TextAnswers = append(em.TextAnswers, &modelTextAnswer)
	}
}

func TestTransformApiSurveyToUtility(t *testing.T) {

	apiModels := &api.Survey{
		Id:    "123",
		Ref:   "/testRef",
		Title: "Test Request",
		Questions: []*api.Question{
			{Id: "1", QuestionnaireId: "1", Text: "How are you?", Order: "1", Kind: "1"},
		},
		RadioPossibleAnswers: []*api.RadioPossibleAnswer{
			{Id: "1", Text: "Great", QuestionId: "1"},
		},
		CheckboxPossibleAnswers: []*api.CheckboxPossibleAnswer{
			{Id: "1", Text: "Great", QuestionId: "1"},
		},
		TextPossibleAnswers: []*api.TextPossibleAnswer{
			{Id: "1", Text: "Great", QuestionId: "1", Placeholder: "Ответь здесь..."},
		},
		RadioAnswers: []*api.RadioAnswer{
			{Id: "1", RadioPossibleAnswerId: "1"},
		},
		CheckboxAnswers: []*api.CheckboxAnswer{
			{Id: "1", CheckboxPossibleAnswerId: "1"},
		},
		TextAnswers: []*api.TextAnswer{
			{Id: "1", TextPossibleAnswerId: "1", Answer: "Bad"},
		},
	}

	var got UtilitySurvey
	var want UtilitySurvey
	got.From(apiModels)

	if got.Questionnaire.Ref != want.Questionnaire.Ref {
		t.Errorf("got %+v\\n, wanted %+v\\n", got, want)
	}

}
