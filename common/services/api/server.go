package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	//"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type GINServer struct {
	svr    *http.Server
	e      *gin.Engine
	config Config
}

func NewGINServer(c Config) *GINServer {
	c.setDefault()

	svr := new(http.Server)
	e := gin.Default()

	if c.Telemetry.Enable {
		//e.Use(otelgin.Middleware(c.Name))
	}

	svr.Addr = c.Addr
	svr.ReadHeaderTimeout = time.Duration(c.Timeout) * time.Millisecond
	svr.ReadTimeout = time.Duration(c.Timeout) * time.Millisecond
	svr.WriteTimeout = time.Duration(c.Timeout) * time.Millisecond
	svr.Handler = e
	return &GINServer{
		svr:    svr,
		e:      e,
		config: c,
	}
}

func (s *GINServer) RegisterRouter(register func(router *gin.Engine)) {
	register(s.e)
}

func (s *GINServer) Run() error {
	c := make(chan error)
	go s.run(c)

	select {
	case err := <-c: //start failed
		return err
	case <-time.After(time.Second): //start success
		//defer s.unregistry()
		//s.registry()
		return <-c //wait exit
	}
}

func (s *GINServer) run(c chan error) {
	defer close(c)
	logrus.Infof("api server starting on %s", s.svr.Addr)
	c <- s.svr.ListenAndServe()
}

//func (s *GINServer) registry() {
//	if !s.config.Registry.Enable {
//		return
//	}
//	registry.Registry(context.Background(), s.config.Registry.Driver, s.config.Name, s.config.Addr)
//}
//
//func (s *GINServer) unregistry() {
//	if !s.config.Registry.Enable {
//		return
//	}
//	//TODO
//}

func (s *GINServer) Shutdown() error {
	return s.svr.Shutdown(context.Background())
}
