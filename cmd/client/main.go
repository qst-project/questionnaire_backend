package main

import (
	"context"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	"google.golang.org/grpc"
	"log"
)

// simple grpc client simple to test server

func main(){
	//var opts []grpc.DialOption
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewQuestionnaireClient(conn)
	res, err := c.Test(context.Background(), &api.TestRequest{})
	res2, err2 := c.GetSurvey(context.Background(), &api.GetSurveyRequest{Ref: "kfmdkfm/ffd"})
	
	if err != nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err)
	}
	log.Println((res.GetResult()))
	log.Println((res2.GetSurvey()))
}
