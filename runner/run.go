package runner

import (
	"context"
	"os"
	"os/signal"
)

var runners = new(Runners)

func Join(r ...Runner) {
	runners.runners = append(runners.runners, r...)
}

func Run(onError func(e error)) {
	for err := range runners.Run() {
		onError(err)
	}
}

func Shutdown(ctx context.Context, onError func(e error)) {
	for err := range runners.Shutdown(ctx) {
		onError(err)
	}
}

func WaitSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
