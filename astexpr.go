package template

import "io"

type astExpr struct {
	operator           string
	operand1, operand2 iAstNode
}

func (n *astExpr) GetImports() []string {
	var res []string

	if n.operand1 != nil {
		res = append(res, n.operand1.GetImports()...)
	}
	if n.operand2 != nil {
		res = append(res, n.operand2.GetImports()...)
	}

	return res
}

func (n *astExpr) GetStrings() []string {
	var res []string

	if n.operand1 != nil {
		res = append(res, n.operand1.GetStrings()...)
	}
	if n.operand2 != nil {
		res = append(res, n.operand2.GetStrings()...)
	}

	return res
}

func (n *astExpr) WriteGo(w io.Writer, opts *GenGoOpts) {
	if n.operand1 != nil {
		n.operand1.WriteGo(w, opts)
	}
	io.WriteString(w, n.operator)

	if n.operand2 != nil {
		n.operand2.WriteGo(w, opts)
	}
}
