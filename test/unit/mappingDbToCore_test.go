package unit

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
	"testing"
)

func TestMappingDbQuestionnaireToCore(t *testing.T) {
	dbQuestionnaire := &gateway.QuestionnaireDB{
		Ref:   "/testRef",
		Title: "Test Request",
	}
	dbQuestions := []*gateway.QuestionDB{
		{
			Type:      1,
			Statement: "Как дела?",
			Order:     1,
		},
	}
	dbOptions := []*gateway.OptionsDB{
		{
			QuestionId: 1,
			Text:       "Хорошо",
		},
		{
			QuestionId: 1,
			Text:       "Нормально",
		},
		{
			QuestionId: 1,
			Text:       "Плохо",
		},
	}

	coreQuestionnaire := dbQuestionnaire.ToCore(dbQuestions, dbOptions)

	if dbQuestionnaire.Ref != coreQuestionnaire.Ref {
		t.Errorf("got %+v\\n, wanted %+v\\n", dbQuestionnaire.Ref, coreQuestionnaire.Ref)
	}
	if dbQuestionnaire.Title != coreQuestionnaire.Title {
		t.Errorf("got %+v\\n, wanted %+v\\n", dbQuestionnaire.Title, coreQuestionnaire.Title)
	}
	if len(coreQuestionnaire.Questions) != len(dbQuestions) {
		t.Errorf("got %+v\\n, wanted %+v\\n", len(coreQuestionnaire.Questions), len(dbQuestions))
	}
	if len(coreQuestionnaire.Questions[0].Options) != len(dbOptions) {
		t.Errorf("got %+v\\n, wanted %+v\\n", len(coreQuestionnaire.Questions[0].Options), len(dbOptions))
	}
}

func TestMappingDbQuestionnaireFromCore(t *testing.T) {
	coreQuestionnaire := &core.Questionnaire{
		Ref:   "/testRef",
		Title: "Test Request",
		Questions: []*core.Question{
			{
				Id:        1,
				Statement: "Как дела?",
				Type:      0,
				Options: []string{
					"Хорошо",
					"Нормально",
					"Плохо",
				}},
		},
	}
	var dbQuestionnaire gateway.QuestionnaireDB
	dbQuestionnaire.FromCore(coreQuestionnaire)

	if dbQuestionnaire.Ref != coreQuestionnaire.Ref {
		t.Errorf("got %+v\\n, wanted %+v\\n", dbQuestionnaire.Ref, coreQuestionnaire.Ref)
	}
	if dbQuestionnaire.Title != coreQuestionnaire.Title {
		t.Errorf("got %+v\\n, wanted %+v\\n", dbQuestionnaire.Title, coreQuestionnaire.Title)
	}
}

func TestMappingDbQuestionToCore(t *testing.T) {
	dbQuestion := &gateway.QuestionDB{
		Statement: "Как дела?",
		Type:      0,
		Order:     1,
	}

	dbOptions := []*gateway.OptionsDB{
		{
			QuestionId: 1,
			Text:       "Хорошо",
		},
		{
			QuestionId: 1,
			Text:       "Нормально",
		},
		{
			QuestionId: 1,
			Text:       "Плохо",
		},
	}

	coreQuestion := dbQuestion.ToCore(dbOptions)

	if coreQuestion.Statement != dbQuestion.Statement {
		t.Errorf("got %+v\\n, wanted %+v\\n", coreQuestion.Statement, dbQuestion.Statement)
	}
	if uint(coreQuestion.Type) != dbQuestion.Type {
		t.Errorf("got %+v\\n, wanted %+v\\n", coreQuestion.Type, dbQuestion.Type)
	}
	if len(coreQuestion.Options) != len(dbOptions) {
		t.Errorf("got %+v\\n, wanted %+v\\n", len(coreQuestion.Options), len(dbOptions))
	}
}

func TestMappingDbQuestionFromCore(t *testing.T) {
	coreQuestion := &core.Question{
		Statement: "Как дела?",
		Type:      0,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}
	var dbQuestion gateway.QuestionDB
	dbQuestion.FromCore(1, 1, coreQuestion)

	if dbQuestion.Statement != coreQuestion.Statement {
		t.Errorf("got %+v\\n, wanted %+v\\n", dbQuestion.Statement, coreQuestion.Statement)
	}
	if dbQuestion.Type != uint(coreQuestion.Type) {
		t.Errorf("got %+v\\n, wanted %+v\\n", dbQuestion.Type, coreQuestion.Type)
	}
}
