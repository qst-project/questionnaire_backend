package delegate

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
)

type QuestionnaireService struct {
	repo gateway.QuestionnaireGateway
}

func NewDelegate(repo *gateway.Module) *QuestionnaireService {
	return &QuestionnaireService{repo: repo}
}

func (s *QuestionnaireService) GetQuestionnaire(ref string) (*api.Questionnaire, error) {
	utilityQuestionnaire, err := s.repo.GetQuestionnaire(ref)
	if err != nil {
		return nil, err
	}
	apiQuestionnaire := utilityQuestionnaire.GetgRPCModel()
	return &apiQuestionnaire, nil
}

func (s *QuestionnaireService) CreateQuestionnaire(Questionnaire *api.Questionnaire) (bool, string, error) {
	var utilityQuestionnaire core.Questionnaire
	utilityQuestionnaire.From(Questionnaire)
	result, err := s.repo.CreateQuestionnaire(utilityQuestionnaire)
	if err != nil {
		return false, "", err
	}
	return result, "generate ref", nil
}

func (s *QuestionnaireService) DeleteQuestionnaire(ref string) (bool, error) {
	result, err := s.repo.DeleteQuestionnaire(ref)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (s *QuestionnaireService) UpdateQuestionnaire(Questionnaire *api.Questionnaire) (bool, error) {
	var utilityQuestionnaire core.Questionnaire
	utilityQuestionnaire.From(Questionnaire)
	result, err := s.repo.UpdateQuestionnaire(utilityQuestionnaire)
	return result, err
}
