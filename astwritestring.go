package template

import "io"

type astWriteString struct {
	value iAstNode
}

func (n *astWriteString) GetImports() []string {
	return append([]string{"io"}, n.value.GetImports()...)
}

func (n *astWriteString) GetStrings() []string {
	return n.value.GetStrings()
}

func (n *astWriteString) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "io.WriteString(w, ")
	n.value.WriteGo(w, opts)
	io.WriteString(w, ")\n")
}
