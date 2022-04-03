package test

import (
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/delegate"
	"github.com/qst-project/backend.git/pkg/gateway"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
)

var lis *bufconn.Listener

const bufSize = 1024 * 1024

func NewTestConfig() pkg.Config {
	return pkg.Config{
		PostgresDsn: "host=localhost port=5432 user=qst password=qst_pwd dbname=qst_db",
	}
}

func init() {
	app := fx.New(
		fx.Provide(pkg.NewLogger()),
		fx.Provide(NewTestConfig),
		fx.Provide(gateway.NewPostgresClient),
		fx.Provide(gateway.NewGatewayModule),
		fx.Provide(delegate.NewDelegateModule),
		fx.Provide(api.NewGrpcHandler),
		fx.Invoke(api.RegisterGrpcServer),
	)
	lis = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	logger := pkg.NewLogger()
	config, _ := pkg.NewConfig()
	postgresClient, _ := gateway.NewPostgresClient(config, logger)
	repos := gateway.NewGatewayModule(postgresClient)
	//serv := delegate.NewDelegateModule(repos)
	//handler, _ := grpc_handler.NewGrpcHandler(serv)
	//api.RegisterQuestionnaireServer(server, handler)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Printf("%v%v", repos, app)
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}
