package template

import "io"

type astFunc struct {
	name   string
	params *astList
}

func (n *astFunc) GetImports() []string {
	if n.params != nil {
		return n.params.GetImports()
	}

	return []string{}
}

func (n *astFunc) GetStrings() []string {
	if n.params != nil {
		return n.params.GetStrings()
	}

	return []string{}
}

func (n *astFunc) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, n.name+"(")
	if n.params != nil {
		for i, param := range n.params.children {
			if i > 0 {
				io.WriteString(w, ", ")
			}
			param.WriteGo(w, opts)
		}
	}
	io.WriteString(w, ")")
}
