package core

type Id uint

type Questionnaire struct {
	Id
	Title string
	Ref   string

	Questions               []Question
	RadioPossibleAnswers    []RadioPossibleAnswer
	CheckboxPossibleAnswers []CheckboxPossibleAnswer
	TextPossibleAnswers     []TextPossibleAnswer
	//RadioAnswers            []*gateway.RadioAnswer
	//CheckboxAnswers         []*gateway.CheckboxAnswer
	//TextAnswers             []*gateway.TextAnswer
}

type Question struct {
	Id
	Order string
	Kind  string
	Text  string
}

type RadioPossibleAnswer struct {
	Id
	QuestionId uint
	Text       string
}

type CheckboxPossibleAnswer struct {
	Id
	QuestionId uint
	Text       string
}

type TextPossibleAnswer struct {
	Id
	QuestionId  uint
	Text        string
	Placeholder string
}

type RadioAnswer struct {
	Id
	QuestionnaireId       uint
	QuestionId            uint
	RadioPossibleAnswerId uint
}

type CheckboxAnswer struct {
	Id
	QuestionnaireId          uint
	QuestionId               uint
	CheckboxPossibleAnswerId uint
}

type TextAnswer struct {
	Id
	QuestionnaireId      uint
	QuestionId           uint
	TextPossibleAnswerId uint
	Answer               string
}
