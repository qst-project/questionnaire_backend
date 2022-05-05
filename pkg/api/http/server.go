package http

import (
	"context"
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/api/http/handlers"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(lifecycle fx.Lifecycle, config pkg.Config, handler handlers.RequestHandler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				server := &http.Server{
					Addr:           config.HttpPort,
					Handler:        cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowCredentials: true}).Handler(handler.InitRoutes()),
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1 << 20,
				}
				go func() {
					if err := server.ListenAndServe(); err != nil {
						log.Fatalf("Failed to listen and serve", err)
					}
				}()
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})
}
