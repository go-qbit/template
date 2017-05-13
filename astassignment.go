package template

import "io"

type astAssignment struct {
	operator             string
	operands1, operands2 *astList
}

func (n *astAssignment) GetImports() []string {
	res := []string{}

	res = append(res, n.operands1.GetImports()...)
	res = append(res, n.operands2.GetImports()...)

	return res
}

func (n *astAssignment) GetStrings() []string {
	res := []string{}

	res = append(res, n.operands1.GetStrings()...)
	res = append(res, n.operands2.GetStrings()...)

	return res
}

func (n *astAssignment) WriteGo(w io.Writer, opts *GenGoOpts) {
	first := true
	for _, child := range n.operands1.children {
		if child != nil {
			if first {
				first = false
			} else {
				io.WriteString(w, ", ")
			}
			child.WriteGo(w, opts)
		}
	}

	io.WriteString(w, n.operator)

	first = true
	for _, child := range n.operands2.children {
		if child != nil {
			if first {
				first = false
			} else {
				io.WriteString(w, ", ")
			}
			child.WriteGo(w, opts)
		}
	}

	io.WriteString(w, "\n")
}
