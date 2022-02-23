package pkg

import (
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/configs"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/grpc"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/repository"
	"go.uber.org/fx"
)

func RunApp() {
	app := fx.New(
		fx.Provide(NewLogger),
		fx.Provide(configs.NewConfig),
		fx.Provide(repository.NewPostgresClient),
		fx.Provide(grpc.NewGrpcHandler),
		fx.Invoke(grpc.RegisterGrpcServer),
	)
	app.Run()
}
