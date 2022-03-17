package test

import (
	"context"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	. "github.com/skinnykaen/quesionnaire_backend.git/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	api.RegisterQuestionnaireServer(server, &RequestHandler{})
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestFirstEndpoint(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.Test(ctx, &api.TestRequest{})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}

func TestSecondEndpoint(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.GetSurvey(ctx, &api.SurveyRequest{Ref: "1"})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}
