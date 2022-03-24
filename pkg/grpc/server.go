package grpc

import (
	"context"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/configs"
	// "github.com/skinnykaen/quesionnaire_backend.git/pkg/repository"
	// "github.com/skinnykaen/quesionnaire_backend.git/pkg/service"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func RegisterGrpcServer(
	lifecycle fx.Lifecycle, handler Handler, config configs.Config,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) (err error) {
				listener, err := net.Listen("tcp", config.GrpcTcpPort)
				if err != nil {
					return
				}
				server := grpc.NewServer()
				api.RegisterQuestionnaireServer(server, handler)
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
