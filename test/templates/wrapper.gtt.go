package templates

import (
	"context"
	"github.com/go-qbit/template/filter"
	"io"
)

func Wrapperpage(ctx context.Context, w io.Writer, tplClbF func(), caption string) {
	io.WriteString(w, "<!DOCTYPE html>\n<html>\n    <head>\n        <title>")
	io.WriteString(w, filter.HTML(caption))
	io.WriteString(w, "</title>\n    </head>\n    <body>")
	tplClbF()
	io.WriteString(w, "</<body>\n</html>")
}
