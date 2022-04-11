package grpc

import (
	"context"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/configs"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
	"time"
)

func RegisterGrpcServer(
	lifecycle fx.Lifecycle, handler Handler, config configs.Config,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) (err error) {
				listener, err := net.Listen("tcp", config.GrpcTcpPort)
				if err != nil {
					log.Fatalln(err)
					return
				}
				server, err := GenerateTLSApi(config.PemPath, config.KeyPath)
				if err != nil {
					log.Fatalln(err)
					return
				}
				api.RegisterQuestionnaireServer(server, handler)

				go func() {
					log.Fatal(server.Serve(listener))
				}()

				grpcWebServer := grpcweb.WrapServer(server)
				multiplex := grpcMultiplexer{
					grpcWebServer,
				}
				r := http.NewServeMux()
				webapp := http.FileServer(http.Dir("../../../Desktop/development/React/frontend"))
				r.Handle("/", multiplex.Handler(webapp))
				srv := &http.Server{
					Handler:      r,
					Addr:         "localhost:8080",
					WriteTimeout: 15 * time.Second,
					ReadTimeout:  15 * time.Second,
				}
				log.Fatal(srv.ListenAndServeTLS(config.PemPath, config.KeyPath))
				//if err = server.Serve(listener); err != nil {
				//	return
				//}
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}

func GenerateTLSApi(pemPath, keyPath string) (*grpc.Server, error) {
	cred, err := credentials.NewServerTLSFromFile(pemPath, keyPath)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer(
		grpc.Creds(cred),
	)
	return s, nil
}

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
