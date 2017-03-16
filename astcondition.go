package template

import "io"

type astCondition struct {
	condition iAstNode
	ifBody    iAstNode
	elseBody  iAstNode
}

func (n *astCondition) GetImports() []string {
	res := []string{}

	if n.condition != nil {
		res = append(res, n.condition.GetImports()...)
	}

	if n.ifBody != nil {
		res = append(res, n.ifBody.GetImports()...)
	}

	if n.elseBody != nil {
		res = append(res, n.elseBody.GetImports()...)
	}

	return res
}

func (n *astCondition) GetStrings() []string {
	res := []string{}

	if n.condition != nil {
		res = append(res, n.condition.GetStrings()...)
	}

	if n.ifBody != nil {
		res = append(res, n.ifBody.GetStrings()...)
	}

	if n.elseBody != nil {
		res = append(res, n.elseBody.GetStrings()...)
	}

	return res
}

func (n *astCondition) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "if ")
	n.condition.WriteGo(w, opts)
	io.WriteString(w, " {\n")
	n.ifBody.WriteGo(w, opts)
	io.WriteString(w, "}")

	if n.elseBody != nil {
		io.WriteString(w, " else {\n")
		n.elseBody.WriteGo(w, opts)
		io.WriteString(w, "}")
	}

	io.WriteString(w, "\n")
}
