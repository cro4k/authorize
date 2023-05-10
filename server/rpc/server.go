package rpc

import (
	"google.golang.org/grpc"

	"github.com/cro4k/authorize/common/services/rpc"
	"github.com/cro4k/authorize/rpc/authrpc"
	"github.com/cro4k/authorize/server/rpc/auth"
)

func NewServer(c rpc.Config) *rpc.GRPCServer {
	svr := rpc.NewGRPCServer(c)
	svr.Register(func(s *grpc.Server) {
		authrpc.RegisterAuthServiceServer(s, auth.NewAuthService())
	})
	return svr
}
