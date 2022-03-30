package pkg

import (
	"github.com/qst-project/backend.git/pkg/configs"
	"github.com/qst-project/backend.git/pkg/grpc"
	"github.com/qst-project/backend.git/pkg/repository"
	"github.com/qst-project/backend.git/pkg/service"
	"go.uber.org/fx"
)

func RunApp() {
	app := fx.New(
		fx.Provide(NewLogger),
		fx.Provide(configs.NewConfig),
		fx.Provide(repository.NewPostgresClient),
		fx.Provide(repository.NewRepository),
		fx.Provide(service.NewService),
		fx.Provide(grpc.NewGrpcHandler),
		fx.Invoke(grpc.RegisterGrpcServer),
	)
	app.Run()
}
