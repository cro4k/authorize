package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/cro4k/authorize/doc/docserver"
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
	defer srv.Shutdown()
	go srv.Run()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
