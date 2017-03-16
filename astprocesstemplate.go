package template

import "io"

type astProcessTemplate struct {
	pkgName string
	name    string
	params  *astList
}

func (n *astProcessTemplate) GetImports() []string {
	if n.params != nil {
		return n.params.GetImports()
	}

	return []string{}
}

func (n *astProcessTemplate) GetStrings() []string {
	if n.params != nil {
		return n.params.GetStrings()
	}

	return []string{}
}

func (n *astProcessTemplate) WriteGo(w io.Writer, opts *GenGoOpts) {
	if n.pkgName != "" {
		io.WriteString(w, n.pkgName+".")
	}
	io.WriteString(w, "Process"+n.name+"(ctx, w")
	if n.params != nil {
		for _, param := range n.params.children {
			io.WriteString(w, ", ")
			param.WriteGo(w, opts)
		}
	}
	io.WriteString(w, ")\n")
}
