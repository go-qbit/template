package template

import "io"

type astParenthesis struct {
	expr iAstNode
}

func (n *astParenthesis) GetImports() []string {
	return n.expr.GetImports()
}

func (n *astParenthesis) GetStrings() []string {
	return n.expr.GetStrings()
}

func (n *astParenthesis) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "(")
	n.expr.WriteGo(w, opts)
	io.WriteString(w, ")")
}
