package api

//func (em *core.Module) From(gRRCModel *Module) {
//
//	em.Questionnaire.From(gRRCModel)
//
//	for _, question := range gRRCModel.GetQuestions() {
//		var modelsQuestion Question
//		modelsQuestion.From(question)
//		em.Questions = append(em.Questions, &modelsQuestion)
//	}
//
//	for _, radioPossibleAnswer := range gRRCModel.GetRadioPossibleAnswers() {
//		var modelRadioPossibleAnswer RadioPossibleAnswer
//		modelRadioPossibleAnswer.From(radioPossibleAnswer)
//		em.RadioPossibleAnswers = append(em.RadioPossibleAnswers, &modelRadioPossibleAnswer)
//	}
//	for _, checkboxPossibleAnswer := range gRRCModel.GetCheckboxPossibleAnswers() {
//		var modelCheckboxPossibleAnswer CheckboxPossibleAnswer
//		modelCheckboxPossibleAnswer.From(checkboxPossibleAnswer)
//		em.CheckboxPossibleAnswers = append(em.CheckboxPossibleAnswers, &modelCheckboxPossibleAnswer)
//	}
//	for _, textPossibleAnswer := range gRRCModel.GetTextPossibleAnswers() {
//		var modelTextPossibleAnswer TextPossibleAnswer
//		modelTextPossibleAnswer.From(textPossibleAnswer)
//		em.TextPossibleAnswers = append(em.TextPossibleAnswers, &modelTextPossibleAnswer)
//	}

//for _, radioAnswer := range gRRCModel.GetRadioAnswers() {
//	var modelRadioAnswer RadioAnswer
//	modelRadioAnswer.From(radioAnswer)
//	em.RadioAnswers = append(em.RadioAnswers, &modelRadioAnswer)
//}
//for _, checkboxAnswer := range gRRCModel.GetCheckboxAnswers() {
//	var modelCheckboxAnswer CheckboxAnswer
//	modelCheckboxAnswer.From(checkboxAnswer)
//	em.CheckboxAnswers = append(em.CheckboxAnswers, &modelCheckboxAnswer)
//}
//for _, textAnswer := range gRRCModel.GetTextAnswers() {
//	var modelTextAnswer TextAnswer
//	modelTextAnswer.From(textAnswer)
//	em.TextAnswers = append(em.TextAnswers, &modelTextAnswer)
//}
//}

//func (em *core.Module) GetgRPCModel() api.Module {
//	var apiQuestionnaire api.Module
//
//	apiQuestionnaire.Id = strconv.FormatUint(uint64(em.Questionnaire.ID), 10)
//	apiQuestionnaire.Title = em.Questionnaire.Title
//	apiQuestionnaire.Ref = em.Questionnaire.Ref
//
//	for _, modelQuestion := range em.Questions {
//		apiQuestion := modelQuestion.GetgRPCModel()
//		apiQuestionnaire.Questions = append(apiQuestionnaire.Questions, &apiQuestion)
//	}
//
//	for _, modelRadioPossibleAnswer := range em.RadioPossibleAnswers {
//		apiRadioPossibleAnswer := modelRadioPossibleAnswer.GetgRPCModel()
//		apiQuestionnaire.RadioPossibleAnswers = append(apiQuestionnaire.RadioPossibleAnswers, &apiRadioPossibleAnswer)
//	}
//	for _, modelCheckboxPossibleAnswer := range em.CheckboxPossibleAnswers {
//		apiCheckboxPossibleAnswer := modelCheckboxPossibleAnswer.GetgRPCModel()
//		apiQuestionnaire.CheckboxPossibleAnswers = append(apiQuestionnaire.CheckboxPossibleAnswers, &apiCheckboxPossibleAnswer)
//	}
//	for _, modelTextPossibleAnswer := range em.TextPossibleAnswers {
//		apiTextPossibleAnswer := modelTextPossibleAnswer.GetgRPCModel()
//		apiQuestionnaire.TextPossibleAnswers = append(apiQuestionnaire.TextPossibleAnswers, &apiTextPossibleAnswer)
//	}

//for _, modelRadioAnswer := range em.RadioAnswers {
//	apiRadioAnswer := modelRadioAnswer.GetGRPCModel()
//	apiQuestionnaire.RadioAnswers = append(apiQuestionnaire.RadioAnswers, &apiRadioAnswer)
//}
//for _, modelCheckboxAnswer := range em.CheckboxAnswers {
//	apiCheckboxAnswer := modelCheckboxAnswer.GetGRPCModel()
//	apiQuestionnaire.CheckboxAnswers = append(apiQuestionnaire.CheckboxAnswers, &apiCheckboxAnswer)
//}
//for _, modelTextAnswer := range em.TextAnswers {
//	apiTextAnswer := modelTextAnswer.GetGRPCModel()
//	apiQuestionnaire.TextAnswers = append(apiQuestionnaire.TextAnswers, &apiTextAnswer)
//}
//
//	return apiQuestionnaire
//}
