package docrouter

import (
	"github.com/cro4k/authorize/doc/annotation"
	"github.com/cro4k/authorize/doc/temporary"
	"github.com/cro4k/authorize/server/api"
	"github.com/cro4k/doc/docer"
	"github.com/cro4k/doc/export/markdown"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Doc(output string, e *gin.Engine) error {
	_ = api.NewServer()
	docer.Init(annotation.Elements())
	documents := docer.Decode(temporary.Get())
	if err := markdown.Export(output, documents.Group(), true); err != nil {
		return err
	}
	e.StaticFS("/", http.Dir(output))
	return nil
}
