package httpModels

type Type uint

const (
	RADIO    Type = 1
	CHECKBOX Type = 2
	TEXT     Type = 3
)

type Questionnaire struct {
	Id        uint        `json:"id"`
	Title     string      `json:"title"`
	Ref       string      `json:"ref"`
	Questions []*Question `json:"questions"`
}

type Question struct {
	Id        uint     `json:"id"`
	Statement string   `json:"statement"`
	Type      Type     `json:"type"`
	Options   []string `json:"options"`
}
