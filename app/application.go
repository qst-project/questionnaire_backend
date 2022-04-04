package app

import (
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/delegate"
	"github.com/qst-project/backend.git/pkg/gateway"
	"github.com/qst-project/backend.git/pkg/usecase"
	"go.uber.org/fx"
)

func RunApp() {
	app := fx.New(
		fx.Provide(pkg.NewLogger),
		fx.Provide(pkg.NewConfig),
		fx.Provide(delegate.Setup),
		fx.Provide(usecase.Setup),
		fx.Provide(gateway.Setup),
		fx.Provide(gateway.NewPostgresClient),
		fx.Provide(api.NewGrpcHandler),
		fx.Invoke(api.RegisterGrpcServer),
	)
	app.Run()
}
