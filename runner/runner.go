package runner

import (
	"context"
	"sync"
)

type Runner interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type Runners struct {
	runners []Runner
	wg      *sync.WaitGroup
}

func (r *Runners) Run() <-chan error {
	r.wg = new(sync.WaitGroup)
	r.wg.Add(len(r.runners))
	e := make(chan error, len(r.runners))
	for _, v := range r.runners {
		go r.run(v, r.wg, e)
	}
	return e
}

func (r *Runners) run(run Runner, wg *sync.WaitGroup, e chan<- error) {
	defer wg.Done()
	if err := run.Run(); err != nil {
		e <- err
	}
}

func (r *Runners) shutdown(ctx context.Context, e chan<- error) {
	n := len(r.runners)
	for i := range r.runners {
		if err := r.runners[n-1-i].Shutdown(ctx); err != nil {
			e <- err
		}
	}
	r.wg.Wait()
	close(e)
}

func (r Runners) Shutdown(ctx context.Context) <-chan error {
	ch := make(chan error, len(r.runners))
	r.shutdown(ctx, ch)
	return ch
}
