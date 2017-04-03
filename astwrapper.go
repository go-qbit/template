package template

import "io"

type astWrapper struct {
	name    string
	vars    *astList
	wrapper *astUseWrapper
	body    iAstNode
}

func (n *astWrapper) GetImports() []string {
	res := []string{"context", "io"}
	if n.body != nil {
		res = append(res, n.body.GetImports()...)
	}

	return res
}

func (n *astWrapper) GetStrings() []string {
	res := []string{}
	if n.body != nil {
		res = append(res, n.body.GetStrings()...)
	}

	return res
}

func (n *astWrapper) WriteGo(w io.Writer, opts *GenGoOpts) {
	io.WriteString(w, "func Wrapper"+n.name+"(ctx context.Context, w io.Writer, tplClbF func()")

	if n.vars != nil {
		for _, v := range n.vars.children {
			if v != nil {
				io.WriteString(w, ", ")
				v.WriteGo(w, opts)
			}
		}
	}

	io.WriteString(w, ") {\n")

	if n.wrapper != nil {
		if n.wrapper.pkgName != "" {
			io.WriteString(w, n.wrapper.pkgName+".")
		}
		io.WriteString(w, "Wrapper"+n.wrapper.name+"(ctx, w, func() {\n")
	}

	if n.body != nil {
		n.body.WriteGo(w, opts)
	}

	if n.wrapper != nil {
		io.WriteString(w, "}")
		if n.wrapper.params != nil {
			for _, p := range n.wrapper.params.children {
				io.WriteString(w, ", ")
				p.WriteGo(w, opts)
			}
		}
		io.WriteString(w, ")\n")
	}

	io.WriteString(w, "}\n\n")
}
