package unit

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/core"
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

type UtilityQuestionnaire struct {
	Questionnaire           *core.QuestionnaireDB
	Questions               []*core.QuestionDB
	RadioPossibleAnswers    []*core.RadioPossibleAnswerDB
	CheckboxPossibleAnswers []*core.CheckboxPossibleAnswerDB
	TextPossibleAnswers     []*core.TextPossibleAnswerDB
	RadioAnswers            []*core.RadioAnswerDB
	CheckboxAnswers         []*core.CheckboxAnswerDB
	TextAnswers             []*core.TextAnswerDB
}

func (em *UtilityQuestionnaire) From(gRRCModel *api.Questionnaire) {
	em.Questionnaire.From(gRRCModel)

	for _, question := range gRRCModel.GetQuestions() {
		var modelsQuestion core.QuestionDB
		modelsQuestion.From(question)
		em.Questions = append(em.Questions, &modelsQuestion)
	}

	for _, radioPossibleAnswer := range gRRCModel.GetRadioPossibleAnswers() {
		var modelRadioPossibleAnswer core.RadioPossibleAnswerDB
		modelRadioPossibleAnswer.From(radioPossibleAnswer)
		em.RadioPossibleAnswers = append(em.RadioPossibleAnswers, &modelRadioPossibleAnswer)
	}
	for _, checkboxPossibleAnswer := range gRRCModel.GetCheckboxPossibleAnswers() {
		var modelCheckboxPossibleAnswer core.CheckboxPossibleAnswerDB
		modelCheckboxPossibleAnswer.From(checkboxPossibleAnswer)
		em.CheckboxPossibleAnswers = append(em.CheckboxPossibleAnswers, &modelCheckboxPossibleAnswer)
	}
	for _, textPossibleAnswer := range gRRCModel.GetTextPossibleAnswers() {
		var modelTextPossibleAnswer core.TextPossibleAnswerDB
		modelTextPossibleAnswer.From(textPossibleAnswer)
		em.TextPossibleAnswers = append(em.TextPossibleAnswers, &modelTextPossibleAnswer)
	}
}

func TestTransformApiQuestionnaireToUtility(t *testing.T) {

	apiModels := &api.Questionnaire{
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
	}

	var got UtilityQuestionnaire
	var want UtilityQuestionnaire
	got.From(apiModels)

	if got.Questionnaire.Ref != want.Questionnaire.Ref {
		t.Errorf("got %+v\\n, wanted %+v\\n", got, want)
	}

}
