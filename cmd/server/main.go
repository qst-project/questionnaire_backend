package server

import (
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/questionnaire"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s:= grpc.NewServer()
	srv := &questionnaire.GRPCServer{}
	api.RegisterQuestionnaireServer(s, srv)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}