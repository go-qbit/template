package template

import "io"

type astIndexedVar struct {
	name, index iAstNode
}

func (*astIndexedVar) GetImports() []string {
	return []string{}
}

func (n *astIndexedVar) GetStrings() []string {
	return append(n.name.GetStrings(), n.index.GetStrings()...)
}

func (n *astIndexedVar) WriteGo(w io.Writer, opts *GenGoOpts) {
	n.name.WriteGo(w, opts)
	io.WriteString(w, "[")
	n.index.WriteGo(w, opts)
	io.WriteString(w, "]")
}
