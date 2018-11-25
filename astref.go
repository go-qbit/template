package template

import "io"

type astRef struct {
	expr iAstNode
}

func (*astRef) GetImports() []string {
	return []string{}
}

func (n *astRef) GetStrings() []string {
	return n.expr.GetStrings()
}

func (n *astRef) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "*(")
	n.expr.WriteGo(w, opts)
	io.WriteString(w, ")")
}
