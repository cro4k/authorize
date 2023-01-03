package rpc

import (
	"context"
	"github.com/cro4k/authorize/config"
	"github.com/cro4k/authorize/rpc/authrpc"
	"github.com/cro4k/authorize/server/rpc/auth"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	srv *grpc.Server
}

func NewServer() *server {
	s := &server{}

	return s
}

func (s *server) Run() error {
	s.srv = grpc.NewServer()
	authrpc.RegisterAuthServiceServer(s.srv, auth.NewAuthService())
	addr, err := net.ResolveTCPAddr("tcp", config.RPCAddr)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	return s.srv.Serve(listener)
}

func (s *server) Shutdown(ctx context.Context) error {
	s.srv.Stop()
	return nil
}
