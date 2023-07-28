package async

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type Engine struct {
	queue <-chan Runner // 外部传入只读任务队列

	ch chan Runner //内部任务队列

	ctx    context.Context
	cancel context.CancelFunc

	wg *sync.WaitGroup

	n int

	count uint32
	limit uint32

	onError func(Runner, error)
}

func NewEngine(queue <-chan Runner, options ...Option) *Engine {
	ctx, cancel := context.WithCancel(context.Background())
	eng := &Engine{
		queue:  queue,
		ctx:    ctx,
		cancel: cancel,
		ch:     make(chan Runner),
		wg:     new(sync.WaitGroup),
	}
	apply(eng, options...)
	return eng
}

func (e *Engine) Start() {
	e.run()
}

func (e *Engine) Wait() {
	e.wg.Wait()
}

func (e *Engine) run() {
	for {
		select {
		case <-e.ctx.Done():
			return
		case runner, ok := <-e.queue:
			if !ok {
				return
			}
			e.distribute(runner) //异步执行
		}
	}
}

// 动态创建协程
func (e *Engine) distribute(runner Runner) {
	select {
	case e.ch <- runner: //

	default:
		if e.available() {
			e.wg.Add(1)
			go e.process(runner, e.ch) //新建协程
		} else {
			e.ch <- runner //阻塞
		}
	}
}

// 处理任务
func (e *Engine) process(current Runner, c <-chan Runner) {
	defer e.wg.Done()
	defer e.release()

	if current != nil {
		e.do(current)
	}

	for {
		select {
		case <-e.ctx.Done():
			return
		case runner, ok := <-c:
			if !ok {
				return
			}
			e.do(runner)
		case <-time.After(time.Second): //释放空闲协程
			return
		}
	}
}

func (e *Engine) do(r Runner) {
	if err := r.Run(); err != nil && e.onError != nil {
		e.onError(r, err)
	}
}

// 检查协程数是否达到上限
func (e *Engine) available() bool {
	if e.limit == 0 {
		return true
	}
	if count := atomic.LoadUint32(&e.count); count >= e.limit {
		return false
	} else {
		atomic.AddUint32(&e.count, 1) // count = count + 1
		return true
	}
}

func (e *Engine) release() {
	if e.limit == 0 {
		return
	}
	atomic.AddUint32(&e.count, ^uint32(0)) // count = count - 1
}

func (e *Engine) Close() {
	e.cancel()
}

func (e *Engine) Count() uint32 {
	return atomic.LoadUint32(&e.count)
}
