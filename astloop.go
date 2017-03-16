package template

import "io"

type astLoop struct {
	expr1, expr2, expr3, body iAstNode
}

func (n *astLoop) GetImports() []string {
	res := []string{}

	if n.expr1 != nil {
		res = append(res, n.expr1.GetImports()...)
	}
	if n.expr2 != nil {
		res = append(res, n.expr2.GetImports()...)
	}
	if n.expr3 != nil {
		res = append(res, n.expr3.GetImports()...)
	}

	if n.body != nil {
		res = append(res, n.body.GetImports()...)
	}

	return res
}

func (n *astLoop) GetStrings() []string {
	res := []string{}

	if n.expr1 != nil {
		res = append(res, n.expr1.GetStrings()...)
	}
	if n.expr2 != nil {
		res = append(res, n.expr2.GetStrings()...)
	}
	if n.expr3 != nil {
		res = append(res, n.expr3.GetStrings()...)
	}

	if n.body != nil {
		res = append(res, n.body.GetStrings()...)
	}

	return res
}

func (n *astLoop) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "for ")

	if n.expr1 != nil {
		n.expr1.WriteGo(w, opts)
	}
	io.WriteString(w, ";")

	if n.expr2 != nil {
		n.expr2.WriteGo(w, opts)
	}
	io.WriteString(w, ";")

	if n.expr3 != nil {
		n.expr3.WriteGo(w, opts)
	}

	io.WriteString(w, "{\n")
	n.body.WriteGo(w, opts)
	io.WriteString(w, "}\n")
}
