package test

import (
	"context"
	"github.com/qst-project/backend.git/app"
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/api"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func RegisterTestGrpcServer(
	lifecycle fx.Lifecycle, handler api.Handler, config pkg.Config, logger *log.Logger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) (err error) {
				//var listener = bufconn.Listen(bufSize)
				server := grpc.NewServer()
				api.RegisterQuestionnaireServiceServer(server, handler)
				go func() {
					err := server.Serve(lis)
					if err != nil {
						logger.Panic(err)
					}
				}()
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}

func TestCreateQuestionnaireEndpoint(t *testing.T) {
	app.AppDI = append(app.AppDI, fx.Invoke(RegisterTestGrpcServer))
	tapp := fx.New(app.AppDI...)
	ctx := context.Background()
	err := tapp.Start(ctx)
	if err != nil {
		panic(err)
	}
	//err = tapp.Stop(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireServiceClient(conn)
	questionnaire := api.Questionnaire{
		Ref: "test_ref",
	}
	resp, err := client.CreateQuestionnaire(ctx, &api.CreateQuestionnaireRequest{Questionnaire: &questionnaire})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetRef())
}
