package docserver

import (
	"context"
	"github.com/cro4k/authorize/doc/docrouter"
	"github.com/cro4k/authorize/runner"
	"github.com/gin-gonic/gin"
	"net/http"
)

type server struct {
	srv    *http.Server
	output string
	addr   string
}

func NewServer(output, addr string) runner.Runner {
	s := &server{output: output, addr: addr}
	return s
}

func (s *server) Run() error {
	e := gin.Default()
	if err := docrouter.Doc(s.output, e); err != nil {
		return err
	}
	s.srv = &http.Server{
		Addr:    s.addr,
		Handler: e,
	}
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
