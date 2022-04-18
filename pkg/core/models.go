package core

type Id uint
type Type uint

const (
	RADIO    Type = 1
	CHECKBOX Type = 2
	TEXT     Type = 3
)

type Questionnaire struct {
	Id
	Title string
	Ref   string

	Questions []*Question
}

type Question struct {
	Id
	Statement string
	Type      Type
	Options   []string
}
