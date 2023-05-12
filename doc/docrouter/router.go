package docrouter

import (
	"net/http"

	"github.com/cro4k/doc/docer"
	"github.com/cro4k/doc/export/markdown"
	"github.com/gin-gonic/gin"

	apix "github.com/cro4k/authorize/common/services/api"
	"github.com/cro4k/authorize/doc/annotation"
	"github.com/cro4k/authorize/doc/temporary"
	"github.com/cro4k/authorize/server/api"
)

func Doc(output string, e *gin.Engine) error {
	_ = api.NewServer(apix.Config{
		Name: "doc",
		Addr: ":8090",
	})
	docer.Init(annotation.Elements())
	documents := docer.Decode(temporary.Get())
	if err := markdown.Export(output, documents.Group(), true); err != nil {
		return err
	}
	e.StaticFS("/", http.Dir(output))
	return nil
}
