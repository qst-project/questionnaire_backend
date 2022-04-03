package gateway

import (
	"github.com/qst-project/backend.git/pkg/core"
	// "github.com/qst-project/backend.git/pkg/delegate"
)

type QuestionnaireRepository struct {
	postgresClient *PostgresClient
}

func NewQuestionnaireGateway(postgresClient *PostgresClient) *QuestionnaireRepository {
	return &QuestionnaireRepository{postgresClient: postgresClient}
}

func (r *QuestionnaireRepository) GetQuestionnaire(ref string) (core.Questionnaire, error) {
	utilityQuestionnaire := core.Questionnaire{}

	// to do разделить на функции
	if r.postgresClient.db.Where("ref = ?", ref).First(&utilityQuestionnaire.Questionnaire).Error != nil {
		return utilityQuestionnaire, r.postgresClient.db.Where("ref = ?", ref).First(&utilityQuestionnaire.Questionnaire).Error
	}
	if r.postgresClient.db.Where("questionnaire_id = ?", utilityQuestionnaire.Questionnaire.ID).Find(&utilityQuestionnaire.Questions).Error != nil {
		return utilityQuestionnaire, r.postgresClient.db.Where("questionnaire_id = ?", utilityQuestionnaire.Questionnaire.ID).Find(&utilityQuestionnaire.Questions).Error
	}

	for _, question := range utilityQuestionnaire.Questions {
		var temporaryRadioPossibleAnswers []*core.RadioPossibleAnswerDB
		var temporaryCheckboxPossibleAnswers []*core.CheckboxPossibleAnswerDB
		var temporaryTextPossibleAnswers []*core.TextPossibleAnswerDB

		if r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryRadioPossibleAnswers).Error != nil {
			return utilityQuestionnaire, r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryRadioPossibleAnswers).Error
		}
		if r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryCheckboxPossibleAnswers).Error != nil {
			return utilityQuestionnaire, r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryCheckboxPossibleAnswers).Error
		}
		if r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryTextPossibleAnswers).Error != nil {
			return utilityQuestionnaire, r.postgresClient.db.Where("question_id = ?", question.ID).Find(&temporaryTextPossibleAnswers).Error
		}
		utilityQuestionnaire.RadioPossibleAnswers = append(utilityQuestionnaire.RadioPossibleAnswers, temporaryRadioPossibleAnswers...)
		utilityQuestionnaire.CheckboxPossibleAnswers = append(utilityQuestionnaire.CheckboxPossibleAnswers, temporaryCheckboxPossibleAnswers...)
		utilityQuestionnaire.TextPossibleAnswers = append(utilityQuestionnaire.TextPossibleAnswers, temporaryTextPossibleAnswers...)
	}

	//for _, radioPossibleAnswer := range utilityQuestionnaire.RadioPossibleAnswers {
	//	var temporaryRadioAnswer []*core.RadioAnswerDB
	//	if r.postgresClient.db.Where("radio_possible_answer_id = ?", radioPossibleAnswer.ID).Find(&temporaryRadioAnswer).Error != nil {
	//		return utilityQuestionnaire, r.postgresClient.db.Where("radio_possible_answer_id = ?", radioPossibleAnswer.ID).Find(&temporaryRadioAnswer).Error
	//	}
	//	utilityQuestionnaire.RadioAnswers = append(utilityQuestionnaire.RadioAnswers, temporaryRadioAnswer...)
	//}
	//
	//for _, checkboxPossibleAnswer := range utilityQuestionnaire.CheckboxAnswers {
	//	var temporaryCheckboxAnswer []*core.CheckboxAnswerDB
	//	if r.postgresClient.db.Where("checkbox_possible_answer_id = ?", checkboxPossibleAnswer.ID).Find(&temporaryCheckboxAnswer).Error != nil {
	//		return utilityQuestionnaire, r.postgresClient.db.Where("checkbox_possible_answer_id = ?", checkboxPossibleAnswer.ID).Find(&temporaryCheckboxAnswer).Error
	//	}
	//	utilityQuestionnaire.CheckboxAnswers = append(utilityQuestionnaire.CheckboxAnswers, temporaryCheckboxAnswer...)
	//}
	//
	//for _, textPossibleAnswer := range utilityQuestionnaire.TextAnswers {
	//	var temporaryTextAnswer []*core.TextAnswerDB
	//	if r.postgresClient.db.Where("text_possible_answer_id = ?", textPossibleAnswer.ID).Find(&temporaryTextAnswer).Error != nil {
	//		return utilityQuestionnaire, r.postgresClient.db.Where("text_possible_answer_id = ?", textPossibleAnswer.ID).Find(&temporaryTextAnswer).Error
	//	}
	//	utilityQuestionnaire.TextAnswers = append(utilityQuestionnaire.TextAnswers, temporaryTextAnswer...)
	//}
	return utilityQuestionnaire, nil
}

func (r *QuestionnaireRepository) CreateQuestionnaire(Questionnaire core.Questionnaire) (bool, error) {
	if r.postgresClient.db.Create(&Questionnaire.Questionnaire).Error != nil {
		return false, r.postgresClient.db.Create(&Questionnaire.Questionnaire).Error
	}
	for _, question := range Questionnaire.Questions {
		if r.postgresClient.db.Create(&question).Error != nil {
			return false, r.postgresClient.db.Create(&question).Error
		}
	}
	for _, radioPossibleAnswer := range Questionnaire.RadioPossibleAnswers {
		if r.postgresClient.db.Create(&radioPossibleAnswer).Error != nil {
			return false, r.postgresClient.db.Create(&radioPossibleAnswer).Error
		}
	}
	for _, checkboxPossibleAnswer := range Questionnaire.CheckboxPossibleAnswers {
		if r.postgresClient.db.Create(&checkboxPossibleAnswer).Error != nil {
			return false, r.postgresClient.db.Create(&checkboxPossibleAnswer).Error
		}
	}
	for _, textPossibleAnswer := range Questionnaire.TextPossibleAnswers {
		if r.postgresClient.db.Create(&textPossibleAnswer).Error != nil {
			return false, r.postgresClient.db.Create(&textPossibleAnswer).Error
		}
	}

	//for _, radioAnswer := range Questionnaire.RadioAnswers {
	//	if r.postgresClient.db.Create(radioAnswer).Error != nil {
	//		return false, r.postgresClient.db.Create(radioAnswer).Error
	//	}
	//}
	//
	//for _, checkboxAnswer := range Questionnaire.CheckboxAnswers {
	//	if r.postgresClient.db.Create(&checkboxAnswer).Error != nil {
	//		return false, r.postgresClient.db.Create(&checkboxAnswer).Error
	//	}
	//}
	//
	//for _, textAnswer := range Questionnaire.TextAnswers {
	//	if r.postgresClient.db.Create(&textAnswer).Error != nil {
	//		return false, r.postgresClient.db.Create(&textAnswer).Error
	//	}
	//}
	return true, nil
}

func (r *QuestionnaireRepository) DeleteQuestionnaire(ref string) (bool, error) {
	var deletedQuestionnaire core.QuestionnaireDB
	if r.postgresClient.db.Where("ref = ?", ref).Delete(&deletedQuestionnaire).Error != nil {
		return false, r.postgresClient.db.Where("ref = ?", ref).Delete(&deletedQuestionnaire).Error
	}
	return true, nil
}

func (r *QuestionnaireRepository) UpdateQuestionnaire(utilityQuestionnaire core.Questionnaire) (bool, error) {

	err := r.postgresClient.db.Model(&utilityQuestionnaire.Questionnaire).
		Updates(core.QuestionnaireDB{Ref: utilityQuestionnaire.Questionnaire.Ref, Title: utilityQuestionnaire.Questionnaire.Title}).Error
	if err != nil {
		return false, nil
	}

	for _, question := range utilityQuestionnaire.Questions {
		err := r.postgresClient.db.Model(&question).
			Updates(core.QuestionDB{Order: question.Order, Kind: question.Kind, Text: question.Text}).Error
		if err != nil {
			return false, err
		}
	}

	for _, radioPossibleAnswer := range utilityQuestionnaire.RadioPossibleAnswers {
		err := r.postgresClient.db.Model(&radioPossibleAnswer).
			Updates(core.RadioPossibleAnswerDB{Text: radioPossibleAnswer.Text}).Error
		if err != nil {
			return false, err
		}
	}
	for _, checkboxPossibleAnswer := range utilityQuestionnaire.CheckboxPossibleAnswers {
		err := r.postgresClient.db.Model(&checkboxPossibleAnswer).
			Updates(core.CheckboxPossibleAnswerDB{Text: checkboxPossibleAnswer.Text}).Error
		if err != nil {
			return false, err
		}
	}
	for _, textPossibleAnswer := range utilityQuestionnaire.TextPossibleAnswers {
		err := r.postgresClient.db.Model(&textPossibleAnswer).
			Updates(core.TextPossibleAnswerDB{Text: textPossibleAnswer.Text}).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
