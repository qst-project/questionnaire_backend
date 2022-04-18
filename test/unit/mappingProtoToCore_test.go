package unit

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/core"
	"testing"
)

func TestMappingProtoQuestionnaireToCore(t *testing.T) {
	protoQuestionnaire := &api.Questionnaire{
		Ref:   "/testRef",
		Title: "Test Request",
		Questions: []*api.Question{
			{Id: "1", Statement: "Как дела?", Type: 0, Options: []string{
				"Хорошо",
				"Нормально",
				"Плохо",
			}},
		},
	}
	coreQuestionnaire := protoQuestionnaire.ToCore()

	if protoQuestionnaire.Ref != coreQuestionnaire.Ref {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestionnaire.GetRef(), coreQuestionnaire.Ref)
	}
	if protoQuestionnaire.Title != coreQuestionnaire.Title {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestionnaire.GetTitle(), coreQuestionnaire.Title)
	}
	if protoQuestionnaire.Questions[0].Statement != coreQuestionnaire.Questions[0].Statement {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestionnaire.Questions[0].Statement, coreQuestionnaire.Questions[0].Statement)
	}
}

func TestMappingProtoQuestionnaireFromCore(t *testing.T) {
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
	var protoQuestionnaire api.Questionnaire
	protoQuestionnaire.FromCore(coreQuestionnaire)

	if protoQuestionnaire.Ref != coreQuestionnaire.Ref {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestionnaire.GetRef(), coreQuestionnaire.Ref)
	}
	if protoQuestionnaire.Title != coreQuestionnaire.Title {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestionnaire.GetTitle(), coreQuestionnaire.Title)
	}
	if protoQuestionnaire.Questions[0].Statement != coreQuestionnaire.Questions[0].Statement {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestionnaire.Questions[0].Statement, coreQuestionnaire.Questions[0].Statement)
	}
}

func TestMappingProtoQuestionToCore(t *testing.T) {
	protoQuestion := &api.Question{
		Statement: "Как дела?",
		Type:      0,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}

	coreQuestion := protoQuestion.ToCore()

	if protoQuestion.Statement != coreQuestion.Statement {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestion.GetStatement(), coreQuestion.Statement)
	}
	if uint(protoQuestion.Type) != uint(coreQuestion.Type) {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestion.GetType(), coreQuestion.Type)
	}
	if protoQuestion.Options[0] != coreQuestion.Options[0] {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestion.Options[0], coreQuestion.Options[0])
	}
}

func TestMappingProtoQuestionFromCore(t *testing.T) {
	coreQuestion := &core.Question{
		Statement: "Как дела?",
		Type:      1,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}
	var protoQuestion api.Question
	protoQuestion.FromCore(coreQuestion)

	if protoQuestion.Statement != coreQuestion.Statement {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestion.GetStatement(), coreQuestion.Statement)
	}
	if uint(protoQuestion.Type) != uint(coreQuestion.Type) {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestion.GetType(), coreQuestion.Type)
	}
	if protoQuestion.Options[0] != coreQuestion.Options[0] {
		t.Errorf("got %+v\\n, wanted %+v\\n", protoQuestion.Options[0], coreQuestion.Options[0])
	}
}
