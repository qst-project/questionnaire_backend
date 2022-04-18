package unit

import (
	"github.com/qst-project/backend.git/pkg/proto"
	"github.com/qst-project/backend.git/test/unit/templates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappingProtoQuestionnaireToCore(t *testing.T) {
	protoQuestionnaire := templates.GetProtoQuestionnaire()
	coreQuestionnaire := protoQuestionnaire.ToCore()

	assert.Equal(t, protoQuestionnaire.Ref, coreQuestionnaire.Ref)
	assert.Equal(t, protoQuestionnaire.Title, coreQuestionnaire.Title)
	assert.Equal(t, protoQuestionnaire.Questions[0].Statement, coreQuestionnaire.Questions[0].Statement)
}

func TestMappingProtoQuestionnaireFromCore(t *testing.T) {
	coreQuestionnaire := templates.GetCoreQuestionnaire()
	var protoQuestionnaire proto.Questionnaire
	protoQuestionnaire.FromCore(coreQuestionnaire)

	assert.Equal(t, protoQuestionnaire.Ref, coreQuestionnaire.Ref)
	assert.Equal(t, protoQuestionnaire.Title, coreQuestionnaire.Title)
	assert.Equal(t, protoQuestionnaire.Questions[0].Statement, coreQuestionnaire.Questions[0].Statement)
}

func TestMappingProtoQuestionToCore(t *testing.T) {
	protoQuestion := templates.GetProtoQuestion()
	coreQuestion := protoQuestion.ToCore()

	assert.Equal(t, protoQuestion.Statement, coreQuestion.Statement)
	assert.Equal(t, uint(protoQuestion.Type), uint(coreQuestion.Type))
	assert.Equal(t, protoQuestion.Options[0], coreQuestion.Options[0])
}

func TestMappingProtoQuestionFromCore(t *testing.T) {
	coreQuestion := templates.GetCoreQuestion()
	var protoQuestion proto.Question
	protoQuestion.FromCore(coreQuestion)

	assert.Equal(t, protoQuestion.Statement, coreQuestion.Statement)
	assert.Equal(t, uint(protoQuestion.Type), uint(coreQuestion.Type))
	assert.Equal(t, protoQuestion.Options[0], coreQuestion.Options[0])
}
