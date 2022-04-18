package main

import (
	"context"
	"github.com/qst-project/backend.git/pkg/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewQuestionnaireServiceClient(conn)
	Questionnaire := &proto.Questionnaire{
		Title: "Update Test Title",
	}
	createQuestionnaireRes, createQuestionnaireErr := c.CreateQuestionnaire(context.Background(), &proto.CreateQuestionnaireRequest{Questionnaire: Questionnaire})
	println("%v %v", createQuestionnaireRes, createQuestionnaireErr)
}
