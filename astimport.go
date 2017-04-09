package template

import (
	"io"
)

type astImport struct {
	alias   string
	pkgName string
}

func (*astImport) GetImports() []string { return []string{} }

func (*astImport) GetStrings() []string { return []string{} }

func (n *astImport) WriteGo(w io.Writer, opts *GenGoOpts) {
	if n.alias != "" {
		io.WriteString(w, n.alias+" ")
	}
	io.WriteString(w, n.pkgName+"\n")
}
