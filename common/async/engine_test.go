package async

import (
	"fmt"
	"testing"
	"time"
)

type TestRunner func() error

func (r TestRunner) Run() error {
	return r()
}

func run(n int) TestRunner {
	return func() error {
		fmt.Println(n)
		time.Sleep(time.Second * 5)
		return nil
	}
}

func TestEngine(t *testing.T) {
	ch := make(chan Runner)

	e := NewEngine(ch, WithLimit(4))

	go e.Start()

	for i := 0; i < 5; i++ {
		ch <- run(i + 1)
	}
	e.Wait()
}
