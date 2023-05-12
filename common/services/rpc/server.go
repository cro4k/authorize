package rpc

import (
	"net"
	"time"

	"github.com/sirupsen/logrus"
	//"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server   *grpc.Server
	config   Config
	listener net.Listener
}

func NewGRPCServer(c Config) *GRPCServer {
	var opts []grpc.ServerOption
	if c.Telemetry.Enable {
		//opts = append(opts, grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
		//opts = append(opts, grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()))
	}
	svr := grpc.NewServer(opts...)

	return &GRPCServer{server: svr, config: c}
}

func (s *GRPCServer) Run() error {

	c := make(chan error)
	go s.run(c)

	select {
	case err := <-c: //start failed
		return err
	case <-time.After(time.Second): //start success
		defer s.unregistry()
		s.registry()
		return <-c //wait exit
	}
}

func (s *GRPCServer) Register(f func(s *grpc.Server)) {
	f(s.server)
}

func (s *GRPCServer) run(c chan error) {
	defer close(c)
	addr, err := net.ResolveTCPAddr("tcp", s.config.Addr)
	if err != nil {
		c <- err
	}
	s.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		c <- err
	}

	logrus.Infof("%s starting on %s", s.config.Name, s.config.Addr)
	c <- s.server.Serve(s.listener)
}

func (s *GRPCServer) registry() {
	//TODO
}

func (s *GRPCServer) unregistry() {
	//TODO
}

func (s *GRPCServer) Shutdown() error {
	s.server.Stop()
	return s.listener.Close()
}
