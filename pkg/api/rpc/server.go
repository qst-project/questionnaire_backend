package rpc

import (
	"context"
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/proto"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RegisterGrpcServer(
	lifecycle fx.Lifecycle, handler Handler, config pkg.Config, logger *log.Logger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) (err error) {
				var listener net.Listener
				if listener, err = net.Listen("tcp", config.GrpcTcpPort); err == nil {
					server := grpc.NewServer()
					proto.RegisterQuestionnaireServiceServer(server, handler)
					go func() {
						err := server.Serve(listener)
						if err != nil {
							logger.Panic(err)
						}
					}()
				}
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}
