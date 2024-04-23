package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cro4k/authorize/internal/app"
)

func main() {
	defer app.Shutdown()

	go app.Run()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT)
	<-ch
}
