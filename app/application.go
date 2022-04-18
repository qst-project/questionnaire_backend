package app

import (
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/delegate"
	"github.com/qst-project/backend.git/pkg/gateway"
	"github.com/qst-project/backend.git/pkg/usecase"
	"go.uber.org/fx"
)

// todo move SQL migration to another invoke
func AppInvokeWith(options ...fx.Option) *fx.App {
	var di = []fx.Option{
		fx.Provide(pkg.NewLogger),
		fx.Provide(pkg.NewConfig),
		fx.Provide(delegate.Setup),
		fx.Provide(usecase.Setup),
		fx.Provide(gateway.Setup),
		fx.Provide(gateway.NewPostgresClient),
		fx.Provide(api.NewGrpcHandler),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	AppInvokeWith(fx.Invoke(api.RegisterGrpcServer)).Run()
}
