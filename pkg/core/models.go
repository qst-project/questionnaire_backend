package core

type Id uint

type Questionnaire struct {
	Id
	Title string
	Ref   string

	Questions []*Question
}

type Question struct {
	Id
	Statement string
	Type      string
	Options   []string
}
