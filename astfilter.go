package template

import "io"

type astFilter struct {
	pkg   string
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
	pkg := n.pkg
	if pkg == "" {
		pkg = "filter"
	}

	io.WriteString(w, pkg+".Filter"+n.name+"(")
	n.value.WriteGo(w, opts)
	io.WriteString(w, ")")
}
