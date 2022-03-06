package main

import (
	"context"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/configs"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/questionnaire"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/repository"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}


func main() {
	app := fx.New(
		fx.Provide(NewLogger),
		fx.Provide(configs.NewConfig),
		fx.Provide(repository.NewPostgresClient),
		fx.Provide(questionnaire.New),
		fx.Invoke(registerHooks),
	)
	app.Run()
}


func registerHooks(
	lifecycle fx.Lifecycle, srv questionnaire.Handler,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				listener, err := net.Listen("tcp", "8080")
				if err != nil {
					log.Fatal(err)
				}
				s:= grpc.NewServer()
				api.RegisterQuestionnaireServer(s, srv)

				if err := s.Serve(listener); err != nil {
					log.Fatal(err)
				}

				return nil
			},
			OnStop: func(context.Context) error{
				return nil
			},
		},
	)
}
