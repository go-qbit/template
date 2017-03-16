package template

import "io"

type astWriteValue struct {
	value iAstNode
}

func (n *astWriteValue) GetImports() []string {
	res := append([]string{"io"}, n.value.GetImports()...)

	if _, ok := n.value.(*astString); !ok {
		res = append(res, "github.com/go-qbit/template/utils")
	}

	return res
}

func (n *astWriteValue) GetStrings() []string {
	return n.value.GetStrings()
}

func (n *astWriteValue) WriteGo(w io.Writer, opts *GenGoOpts) {
	switch n.value.(type) {
	case *astString:
		io.WriteString(w, "w.Write(")
		n.value.WriteGo(w, opts)
		io.WriteString(w, ")\n")
	default:
		io.WriteString(w, "io.WriteString(w, utils.ToString(")
		n.value.WriteGo(w, opts)
		io.WriteString(w, "))\n")
	}
}
