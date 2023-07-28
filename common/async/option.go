package async

import "runtime"

type Option func(e *Engine)

func apply(e *Engine, options ...Option) {
	for _, opt := range options {
		opt(e)
	}
}

// WithLimit 限制最大并发数，limit==0 时不限制
func WithLimit(limit uint32) Option {
	return func(e *Engine) {
		e.limit = limit
	}
}

func WithErrorHandler(onError func(runner Runner, err error)) Option {
	return func(e *Engine) {
		e.onError = onError
	}
}

func WithCpuNums(magnify ...uint32) Option {
	var n uint32 = 1
	if len(magnify) > 0 && magnify[0] > 0 {
		n = magnify[0]
	}
	return WithLimit(n * uint32(runtime.NumCPU()))
}
