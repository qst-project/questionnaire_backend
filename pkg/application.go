package pkg

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/delegate"
	"github.com/qst-project/backend.git/pkg/gateway"
	"go.uber.org/fx"
)

func RunApp() {
	app := fx.New(
		fx.Provide(NewLogger),
		fx.Provide(NewConfig),
		fx.Provide(gateway.NewPostgresClient),
		fx.Provide(gateway.NewGatewayModule),
		fx.Provide(delegate.NewDelegateModule),
		fx.Provide(api.NewGrpcHandler),
		fx.Invoke(api.RegisterGrpcServer),
	)
	app.Run()
}
