package template

import "io"

type astTypeConv struct {
	expr  iAstNode
	tType string
}

func (*astTypeConv) GetImports() []string {
	return []string{}
}

func (n *astTypeConv) GetStrings() []string {
	return n.expr.GetStrings()
}

func (n *astTypeConv) WriteGo(w io.Writer, opts *GenGoOpts) {
	n.expr.WriteGo(w, opts)
	io.WriteString(w, ".("+n.tType+")")
}
