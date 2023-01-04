package main

import (
	"context"
	"github.com/cro4k/authorize/config"
	"github.com/cro4k/authorize/doc/docserver"
	_ "github.com/cro4k/authorize/logs"
	"github.com/cro4k/authorize/runner"
	"github.com/cro4k/authorize/server/api"
	"github.com/cro4k/authorize/server/rpc"
)

//go:generate ann build -o doc/annotation
func main() {
	runner.Join(
		api.NewServer(),
		rpc.NewServer(),
	)
	if config.C().Env == config.Develop {
		runner.Join(docserver.NewServer("doc", config.DOCAddr))
	}
	runner.Run(func(e error) {})
	runner.WaitSignal()
	runner.Shutdown(context.Background(), func(e error) {})
}
