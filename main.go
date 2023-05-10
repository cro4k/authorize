package main

import (
	"github.com/cro4k/authorize/common/conf"
	"github.com/cro4k/authorize/common/services/service"
	"github.com/cro4k/authorize/config"
	"github.com/cro4k/authorize/internal/dao"
	"github.com/cro4k/authorize/internal/db"
	_ "github.com/cro4k/authorize/logs"
	"github.com/cro4k/authorize/server/api"
	"github.com/cro4k/authorize/server/rpc"
)

//go:generate ann build -o doc/annotation
func main() {
	var c config.Config
	conf.MustLoad(&c)

	db.Setup(c.DB)
	dao.Migrate(db.DB())

	group := service.Group(
		api.NewServer(c.API),
		rpc.NewServer(c.RPC),
	)

	defer group.Shutdown()
	group.Run()
	service.WaitSignal()
}
