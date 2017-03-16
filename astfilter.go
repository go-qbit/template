package template

import "io"

type astFilter struct {
	name  string
	value iAstNode
}

func (n *astFilter) GetImports() []string {
	return append([]string{"github.com/go-qbit/template/filter"}, n.value.GetImports()...)
}

func (n *astFilter) GetStrings() []string {
	return n.value.GetStrings()
}

func (n *astFilter) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "filter."+n.name+"(")
	n.value.WriteGo(w, opts)
	io.WriteString(w, ")")
}
