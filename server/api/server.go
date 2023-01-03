package api

import (
	"context"
	"github.com/cro4k/authorize/config"
	"github.com/cro4k/authorize/runner"
	"github.com/cro4k/authorize/server/api/router"
	"net/http"
)

type server struct {
	srv *http.Server
}

func NewServer() runner.Runner {
	s := &server{}
	s.srv = &http.Server{
		Addr:    config.APIAddr,
		Handler: router.Router(),
	}
	return s
}

func (s *server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
