package todo

import (
	"context"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *Server {

	var port string = viper.GetString("Port")

	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Run() error {

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {

	return s.httpServer.Shutdown(ctx)
}
