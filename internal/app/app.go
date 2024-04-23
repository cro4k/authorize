package app

import (
	"context"
	"github.com/cro4k/authorize/internal/dao/migrate"

	"github.com/go-chocolate/chocolate/pkg/chocolate/chocohttp"
	"github.com/go-chocolate/chocolate/pkg/chocolate/chocorpc"
	"github.com/go-chocolate/chocolate/pkg/chocolate/service"
	"github.com/go-chocolate/configuration/configuration"
	"github.com/sirupsen/logrus"

	"github.com/cro4k/authorize/internal/app/config"
	"github.com/cro4k/authorize/internal/app/dependency"
	"github.com/cro4k/authorize/internal/entrance/http"
	"github.com/cro4k/authorize/internal/module"
)

var ctx, cancel = context.WithCancel(context.Background())

func Run() {
	var cfg config.Config
	if err := configuration.Load(&cfg); err != nil {
		panic(err)
	}
	if err := dependency.Setup(cfg); err != nil {
		panic(err)
	}
	if err := module.Setup(); err != nil {
		panic(err)
	}

	httpsrv := chocohttp.NewServer(cfg.HTTP)
	httpsrv.SetRouter(http.Router())
	logrus.Infof("http server listening on %s", cfg.HTTP.Addr)

	rpcsrv := chocorpc.NewServer(cfg.RPC)
	logrus.Infof("rpc server listening on %s", cfg.RPC.Addr)

	if err := migrate.Migrate(dependency.Get().DB); err != nil {
		panic(err)
	}

	group := service.Group(httpsrv, rpcsrv)

	if err := group.Run(ctx); err != nil {
		panic(err)
	}

}

func Shutdown() {
	cancel()
}
