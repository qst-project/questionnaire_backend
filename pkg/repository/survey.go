package repository

import (
	"github.com/qst-project/backend.git/pkg/models"
	// "github.com/qst-project/backend.git/pkg/service"
)

type SurveyRepository struct {
	postgresClient *PostgresClient
}

func NewSurveyRepository(postgresClient *PostgresClient) *SurveyRepository {
	return &SurveyRepository{postgresClient: postgresClient}
}

func (r *SurveyRepository) GetSurvey(ref string) (models.UtilitySurvey, error) {
	utilitySurvey := models.UtilitySurvey{}

	// to do разделить на функции
	if r.postgresClient.db.Where("ref = ?", ref).First(&utilitySurvey.Questionnaire).Error != nil {
		return utilitySurvey, r.postgresClient.db.Where("ref = ?", ref).First(&utilitySurvey.Questionnaire).Error
	}
	if r.postgresClient.db.Where("questionnaire_id = ?", utilitySurvey.Questionnaire.ID).Find(&utilitySurvey.Questions).Error != nil {
		return utilitySurvey, r.postgresClient.db.Where("questionnaire_id = ?", utilitySurvey.Questionnaire.ID).Find(&utilitySurvey.Questions).Error
	}

	for _, question := range utilitySurvey.Questions {
		var temporaryRadioPossibleAnswers []*models.RadioPossibleAnswer
		var temporaryCheckboxPossibleAnswers []*models.CheckboxPossibleAnswer
		var temporaryTextPossibleAnswers []*models.TextPossibleAnswer

		if r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryRadioPossibleAnswers).Error != nil {
			return utilitySurvey, r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryRadioPossibleAnswers).Error
		}
		if r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryCheckboxPossibleAnswers).Error != nil {
			return utilitySurvey, r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryCheckboxPossibleAnswers).Error
		}
		if r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryTextPossibleAnswers).Error != nil {
			return utilitySurvey, r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryTextPossibleAnswers).Error
		}
		utilitySurvey.RadioPossibleAnswers = append(utilitySurvey.RadioPossibleAnswers, temporaryRadioPossibleAnswers...)
		utilitySurvey.CheckboxPossibleAnswers = append(utilitySurvey.CheckboxPossibleAnswers, temporaryCheckboxPossibleAnswers...)
		utilitySurvey.TextPossibleAnswers = append(utilitySurvey.TextPossibleAnswers, temporaryTextPossibleAnswers...)
	}

	//for _, radioPossibleAnswer := range utilitySurvey.RadioPossibleAnswers {
	//	var temporaryRadioAnswer []*models.RadioAnswer
	//	if r.postgresClient.db.Where("radio_possible_answer_id = ?", radioPossibleAnswer.ID).Find(&temporaryRadioAnswer).Error != nil {
	//		return utilitySurvey, r.postgresClient.db.Where("radio_possible_answer_id = ?", radioPossibleAnswer.ID).Find(&temporaryRadioAnswer).Error
	//	}
	//	utilitySurvey.RadioAnswers = append(utilitySurvey.RadioAnswers, temporaryRadioAnswer...)
	//}
	//
	//for _, checkboxPossibleAnswer := range utilitySurvey.CheckboxAnswers {
	//	var temporaryCheckboxAnswer []*models.CheckboxAnswer
	//	if r.postgresClient.db.Where("checkbox_possible_answer_id = ?", checkboxPossibleAnswer.ID).Find(&temporaryCheckboxAnswer).Error != nil {
	//		return utilitySurvey, r.postgresClient.db.Where("checkbox_possible_answer_id = ?", checkboxPossibleAnswer.ID).Find(&temporaryCheckboxAnswer).Error
	//	}
	//	utilitySurvey.CheckboxAnswers = append(utilitySurvey.CheckboxAnswers, temporaryCheckboxAnswer...)
	//}
	//
	//for _, textPossibleAnswer := range utilitySurvey.TextAnswers {
	//	var temporaryTextAnswer []*models.TextAnswer
	//	if r.postgresClient.db.Where("text_possible_answer_id = ?", textPossibleAnswer.ID).Find(&temporaryTextAnswer).Error != nil {
	//		return utilitySurvey, r.postgresClient.db.Where("text_possible_answer_id = ?", textPossibleAnswer.ID).Find(&temporaryTextAnswer).Error
	//	}
	//	utilitySurvey.TextAnswers = append(utilitySurvey.TextAnswers, temporaryTextAnswer...)
	//}
	return utilitySurvey, nil
}

func (r *SurveyRepository) SetSurvey(survey models.UtilitySurvey) (bool, error) {
	if r.postgresClient.db.Create(&survey.Questionnaire).Error != nil {
		return false, r.postgresClient.db.Create(&survey.Questionnaire).Error
	}
	for _, question := range survey.Questions {
		if r.postgresClient.db.Create(&question).Error != nil {
			return false, r.postgresClient.db.Create(&question).Error
		}
	}
	for _, radioPossibleAnswer := range survey.RadioPossibleAnswers {
		if r.postgresClient.db.Create(&radioPossibleAnswer).Error != nil {
			return false, r.postgresClient.db.Create(&radioPossibleAnswer).Error
		}
	}
	for _, checkboxPossibleAnswer := range survey.CheckboxPossibleAnswers {
		if r.postgresClient.db.Create(&checkboxPossibleAnswer).Error != nil {
			return false, r.postgresClient.db.Create(&checkboxPossibleAnswer).Error
		}
	}
	for _, textPossibleAnswer := range survey.TextPossibleAnswers {
		if r.postgresClient.db.Create(&textPossibleAnswer).Error != nil {
			return false, r.postgresClient.db.Create(&textPossibleAnswer).Error
		}
	}

	//for _, radioAnswer := range survey.RadioAnswers {
	//	if r.postgresClient.db.Create(radioAnswer).Error != nil {
	//		return false, r.postgresClient.db.Create(radioAnswer).Error
	//	}
	//}
	//
	//for _, checkboxAnswer := range survey.CheckboxAnswers {
	//	if r.postgresClient.db.Create(&checkboxAnswer).Error != nil {
	//		return false, r.postgresClient.db.Create(&checkboxAnswer).Error
	//	}
	//}
	//
	//for _, textAnswer := range survey.TextAnswers {
	//	if r.postgresClient.db.Create(&textAnswer).Error != nil {
	//		return false, r.postgresClient.db.Create(&textAnswer).Error
	//	}
	//}
	return true, nil
}
