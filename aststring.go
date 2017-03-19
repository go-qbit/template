package template

import (
	"io"
	"strings"
)

type astString struct {
	value string
}

func (*astString) GetImports() []string {
	return []string{}
}

func (n *astString) GetStrings() []string {
	return []string{}
}

func (n *astString) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, strings.Replace(n.value, "\n", `\n`, -1))
}
