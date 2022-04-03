package api

import (
	"context"
	"github.com/qst-project/backend.git/pkg"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func RegisterGrpcServer(
	lifecycle fx.Lifecycle, handler Handler, config pkg.Config,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) (err error) {
				listener, err := net.Listen("tcp", config.GrpcTcpPort)
				if err != nil {
					return
				}
				server := grpc.NewServer()
				RegisterQuestionnaireServiceServer(server, handler)
				if err = server.Serve(listener); err != nil {
					return
				}
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}
