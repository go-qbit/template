package template

import "io"

type astAssignment struct {
	operator           string
	operand1, operand2 iAstNode
}

func (n *astAssignment) GetImports() []string {
	res := []string{}

	res = append(res, n.operand1.GetImports()...)
	res = append(res, n.operand2.GetImports()...)

	return res
}

func (n *astAssignment) GetStrings() []string {
	res := []string{}

	res = append(res, n.operand1.GetStrings()...)
	res = append(res, n.operand2.GetStrings()...)

	return res
}

func (n *astAssignment) WriteGo(w io.Writer, opts *GenGoOpts) {
	n.operand1.WriteGo(w, opts)
	io.WriteString(w, n.operator)
	n.operand2.WriteGo(w, opts)
	io.WriteString(w, "\n")
}
