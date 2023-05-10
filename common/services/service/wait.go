package service

import (
	"os"
	"os/signal"
)

func WaitSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
