package main

import (
	"context"
	_ "github.com/cro4k/authorize/logs"
	"github.com/cro4k/authorize/runner"
	"github.com/cro4k/authorize/server/api"
	"github.com/cro4k/authorize/server/rpc"
)

func main() {
	runner.Join(
		api.NewServer(),
		rpc.NewServer(),
	)
	runner.Run(func(e error) {})
	runner.WaitSignal()
	runner.Shutdown(context.Background(), func(e error) {})
}
