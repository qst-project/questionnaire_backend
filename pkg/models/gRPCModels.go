package models

// import (
// 	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
// 	"log"
// 	"strconv"
// )

// func (em *Questionnaire) GetgRPCModel() api.Survey {
// 	return api.Survey{
// 		Id: strconv.FormatUint(uint64(em.ID), 10),
// 		Title: em.Title,
// 	}
// }

// func (em *Questionnaire) From(gRRCModel api.Survey) {
// 	u, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	em.ID = uint(u)
// 	em.Title = gRRCModel.Title
// }

// func (em *Question) GetgRPCModel() api.Question {
// 	return api.Question{
// 		Id: strconv.FormatUint(uint64(em.ID), 10),
// 		Text: em.Title,
// 		Order: em.Order,
// 		Kind: em.Kind,
// 	}
// }

// func (em *Question) From(gRRCModel api.Question) {
// 	u, err := strconv.ParseUint(gRRCModel.Id, 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	em.ID = uint(u)
// 	em.Text = gRRCModel.Text
// 	em.Order = gRRCModel.Order
// 	em.Kind = gRRCModel.Kind
// }
