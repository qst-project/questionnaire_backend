package main

import (
	"context"
	"github.com/qst-project/backend.git/pkg/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewQuestionnaireServiceClient(conn)
	Questionnaire := &api.Questionnaire{
		Title: "Test title",
	}
	response, _ := c.CreateQuestionnaire(context.Background(), &api.CreateQuestionnaireRequest{Questionnaire: Questionnaire})
	println("%v", response.String())
}