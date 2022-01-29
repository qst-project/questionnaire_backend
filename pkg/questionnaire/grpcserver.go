package questionnaire

import (
	"context"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
)

// gRPCServer
type GRPCServer struct {

}

func (s *GRPCServer) Test(ctx context.Context, req *api.TestRequest) (*api.TestResponse, error){
	return & api.TestResponse{Result: "Test 1,2"}, nil
}