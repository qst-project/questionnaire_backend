package test

import (
	"context"
	"github.com/qst-project/backend.git/app"
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/delegate"
	"github.com/qst-project/backend.git/pkg/proto"
	"github.com/qst-project/backend.git/test/unit/templates"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func PanicOnError(t *testing.T, err error) {
	if assert.Error(t, err) {
		panic(err)
	}
}
func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func RegisterTestGrpcServer(
	lifecycle fx.Lifecycle, handler api.Handler, config pkg.Config, logger *log.Logger, questionnaireDelegate delegate.QuestionnaireDelegate,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) (err error) {
				//var listener = bufconn.Listen(bufSize)
				server := grpc.NewServer()
				proto.RegisterQuestionnaireServiceServer(server, handler)
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

func InvokeTestApp(t *testing.T, ctx context.Context) {
	err := app.AppInvokeWith(fx.Invoke(RegisterTestGrpcServer)).Start(ctx)
	assert.NoError(t, err)
}
func InvokeTestClient(t *testing.T, ctx context.Context) (conn *grpc.ClientConn) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err)
	return
}

func CloseConnection(t *testing.T, c *grpc.ClientConn) {
	err := c.Close()
	assert.NoError(t, err)
}

func TestCreateQuestionnaireUnary(t *testing.T) {
	ctx := context.Background()
	InvokeTestApp(t, ctx)
	c := InvokeTestClient(t, ctx)
	defer CloseConnection(t, c)
	client := proto.NewQuestionnaireServiceClient(c)
	questionnaire := templates.GetProtoQuestionnaire()
	resp, err := client.CreateQuestionnaire(ctx, &proto.CreateQuestionnaireRequest{Questionnaire: questionnaire})
	assert.NoError(t, err)
	assert.Equal(t, questionnaire.Ref, resp.Ref)
}

func TestGetQuestionnaireUnary(t *testing.T) {
	ctx := context.Background()
	InvokeTestApp(t, ctx)
	c := InvokeTestClient(t, ctx)
	defer CloseConnection(t, c)
	client := proto.NewQuestionnaireServiceClient(c)
	resp, err := client.GetQuestionnaire(ctx, &proto.GetQuestionnaireRequest{Ref: "testRef"})
	assert.NoError(t, err)
	assert.Equal(t, "Test Request", resp.GetQuestionnaire().Title)
}
