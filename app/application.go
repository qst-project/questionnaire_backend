package app

import (
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/gateway"
	"github.com/qst-project/backend.git/pkg/usecase"
	"go.uber.org/fx"
)

var AppDI = []fx.Option{
	fx.Provide(pkg.NewLogger),
	fx.Provide(pkg.NewConfig),
	fx.Provide(api.Setup),
	fx.Provide(usecase.Setup),
	fx.Provide(gateway.Setup),
	fx.Provide(gateway.NewPostgresClient),
	fx.Provide(api.NewGrpcHandler),
}

func RunApp() {
	AppDI = append(AppDI, fx.Invoke(api.RegisterGrpcServer))
	fx.New(AppDI...).Run()
}
