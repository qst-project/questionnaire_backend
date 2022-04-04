package api

import "github.com/qst-project/backend.git/pkg/core"

func (q *Questionnaire) ToCore() (qst core.Questionnaire, err error) {
	qst = core.Questionnaire{
		Id:    core.Id(q.Id),
		Title: q.Title,
		Ref:   q.Ref,
	}
	return
}
