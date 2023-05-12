package docserver

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cro4k/authorize/doc/docrouter"
)

type DocServer struct {
	srv    *http.Server
	output string
	addr   string
}

func NewServer(output, addr string) *DocServer {
	s := &DocServer{output: output, addr: addr}
	return s
}

func (s *DocServer) Run() error {
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

func (s *DocServer) Shutdown() error {
	return s.srv.Shutdown(context.Background())
}
