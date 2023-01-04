package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cro4k/authorize/doc/docserver"
	"github.com/cro4k/authorize/runner"
)

var (
	output string
	port   int
)

func init() {
	flag.StringVar(&output, "o", "doc/output", "")
	flag.IntVar(&port, "p", 8090, "")
	flag.Parse()
}

func main() {
	srv := docserver.NewServer(output, fmt.Sprintf(":%d", port))
	runner.Join(srv)
	runner.Run(func(e error) {})
	runner.WaitSignal()
	runner.Shutdown(context.Background(), func(e error) {})
}
