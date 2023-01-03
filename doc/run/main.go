package main

import (
	"github.com/cro4k/authorize/doc/annotation"
	"github.com/cro4k/authorize/doc/temporary"
	"github.com/cro4k/authorize/server/api"
	"github.com/cro4k/doc/docer"
	"github.com/cro4k/doc/export/markdown"
	"log"
)

func main() {
	_ = api.NewServer()
	docer.Init(annotation.Elements())
	documents := docer.Decode(temporary.Get())
	err := markdown.Export("doc/output", documents.Group(), true)
	if err != nil {
		log.Println(err)
	}
}
